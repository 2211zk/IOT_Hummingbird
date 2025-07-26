package wl_department

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/response"
)

// 缓存键常量
const (
	CACHE_KEY_DEPARTMENT_TREE    = "department:tree"
	CACHE_KEY_DEPARTMENT_LIST    = "department:list:%s"
	CACHE_KEY_DEPARTMENT_DETAIL  = "department:detail:%d"
	CACHE_KEY_AVAILABLE_DEVICES  = "department:devices:available:%s"
	CACHE_KEY_DEPARTMENT_DEVICES = "department:devices:%d"
	CACHE_EXPIRATION_TIME        = 5 * time.Minute
	CACHE_TREE_EXPIRATION_TIME   = 10 * time.Minute
)

// DepartmentCache 部门缓存服务
type DepartmentCache struct {
	mu    sync.RWMutex
	cache map[string]CacheItem
}

// CacheItem 缓存项
type CacheItem struct {
	Data      interface{}
	ExpiresAt time.Time
}

var departmentCache *DepartmentCache
var cacheOnce sync.Once

// GetDepartmentCache 获取缓存实例（单例模式）
func GetDepartmentCache() *DepartmentCache {
	cacheOnce.Do(func() {
		departmentCache = &DepartmentCache{
			cache: make(map[string]CacheItem),
		}
		// 启动清理协程
		go departmentCache.startCleanup()
	})
	return departmentCache
}

// Get 获取缓存数据
func (c *DepartmentCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.cache[key]
	if !exists || time.Now().After(item.ExpiresAt) {
		return nil, false
	}

	return item.Data, true
}

// Set 设置缓存数据
func (c *DepartmentCache) Set(key string, data interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheItem{
		Data:      data,
		ExpiresAt: time.Now().Add(expiration),
	}
}

// Delete 删除缓存数据
func (c *DepartmentCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.cache, key)
}

// Clear 清空所有缓存
func (c *DepartmentCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]CacheItem)
}

// InvalidateDepartmentCache 使部门相关缓存失效
func (c *DepartmentCache) InvalidateDepartmentCache() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 删除所有部门相关的缓存
	for key := range c.cache {
		if len(key) > 10 && key[:10] == "department" {
			delete(c.cache, key)
		}
	}
}

// startCleanup 启动定期清理过期缓存
func (c *DepartmentCache) startCleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		}
	}
}

// cleanup 清理过期缓存
func (c *DepartmentCache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.cache {
		if now.After(item.ExpiresAt) {
			delete(c.cache, key)
		}
	}
}

// 缓存辅助方法

// GetDepartmentTreeFromCache 从缓存获取部门树
func (s *WlDepartmentService) GetDepartmentTreeFromCache(excludeID *uint) ([]*response.DepartmentTreeNode, error) {
	cache := GetDepartmentCache()
	cacheKey := CACHE_KEY_DEPARTMENT_TREE
	if excludeID != nil {
		cacheKey = fmt.Sprintf("%s:exclude:%d", CACHE_KEY_DEPARTMENT_TREE, *excludeID)
	}

	if cached, exists := cache.Get(cacheKey); exists {
		if tree, ok := cached.([]*response.DepartmentTreeNode); ok {
			return tree, nil
		}
	}

	// 缓存未命中，从数据库获取
	tree, err := s.getDepartmentTreeFromDB(excludeID)
	if err != nil {
		return nil, err
	}

	// 存入缓存
	cache.Set(cacheKey, tree, CACHE_TREE_EXPIRATION_TIME)
	return tree, nil
}

// getDepartmentTreeFromDB 从数据库获取部门树
func (s *WlDepartmentService) getDepartmentTreeFromDB(excludeID *uint) ([]*response.DepartmentTreeNode, error) {
	var departments []wl_department.WlDepartment
	db := global.GVA_DB.Model(&wl_department.WlDepartment{})

	// 如果指定了排除ID，则排除该部门及其所有子部门
	if excludeID != nil {
		excludeIDs, err := s.getAllChildrenIDs(*excludeID)
		if err != nil {
			return nil, err
		}
		excludeIDs = append(excludeIDs, *excludeID)
		db = db.Where("id NOT IN ?", excludeIDs)
	}

	err := db.Where("status = ?", "启用").
		Select("id, parent_id, name, sort, created_at").
		Order("sort asc, created_at desc").Find(&departments).Error
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	deptTree := s.buildDepartmentTreeOptimized(departments, nil)
	var tree []*response.DepartmentTreeNode
	for _, dept := range deptTree {
		tree = append(tree, response.ConvertToTreeNode(&dept))
	}

	return tree, nil
}

// GetDepartmentListFromCache 从缓存获取部门列表
func (s *WlDepartmentService) GetDepartmentListFromCache(req request.WlDepartmentSearch) ([]wl_department.WlDepartment, int64, error) {
	cache := GetDepartmentCache()

	// 生成缓存键
	reqBytes, _ := json.Marshal(req)
	cacheKey := fmt.Sprintf(CACHE_KEY_DEPARTMENT_LIST, string(reqBytes))

	if cached, exists := cache.Get(cacheKey); exists {
		if result, ok := cached.(struct {
			List  []wl_department.WlDepartment
			Total int64
		}); ok {
			return result.List, result.Total, nil
		}
	}

	// 缓存未命中，从数据库获取
	list, total, err := s.GetWlDepartmentList(req)
	if err != nil {
		return nil, 0, err
	}

	// 存入缓存（只缓存查询结果较小的情况）
	if len(list) <= 100 {
		cache.Set(cacheKey, struct {
			List  []wl_department.WlDepartment
			Total int64
		}{List: list, Total: total}, CACHE_EXPIRATION_TIME)
	}

	return list, total, nil
}

// InvalidateCache 使缓存失效（在数据更新时调用）
func (s *WlDepartmentService) InvalidateCache() {
	cache := GetDepartmentCache()
	cache.InvalidateDepartmentCache()
}
