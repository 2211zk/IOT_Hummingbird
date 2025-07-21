package wl_playform

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_playform/request"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type WlResourcesService struct{}

// CreateWlResourcesWithTransaction 创建资源（事务处理）
func (wlResourcesService *WlResourcesService) CreateWlResourcesWithTransaction(ctx context.Context, wlResources *wl_playform.WlResources, resourceData map[string]interface{}) error {
	// 开始MySQL事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 生成MongoDB ID和Key
	mongoID := primitive.NewObjectID()
	mongoKey := time.Now().Format("20060102150405") + "_" + *wlResources.InstanceName
	wlResources.ResourcesKey = &mongoKey

	// 检查MongoDB key是否已存在
	collection := global.GVA_WANLIAN_MONGO.Database.Collection("wl_resources")
	var existingDoc bson.M
	err := collection.Find(ctx, bson.M{"mongo_key": mongoKey}).One(&existingDoc)
	if err == nil {
		// 如果找到了相同key的文档，说明可能重复提交
		global.GVA_LOG.Warn("检测到重复的MongoDB key", zap.String("mongo_key", mongoKey))
		tx.Rollback()
		return errors.New("检测到重复提交，请稍后重试")
	}

	// 创建MongoDB文档
	mongoDoc := bson.M{
		"_id":           mongoID,
		"mongo_key":     mongoKey,
		"instance_name": *wlResources.InstanceName,
		"resource_data": resourceData,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
	}

	// 添加调试日志
	global.GVA_LOG.Info("准备插入MongoDB数据",
		zap.String("mongo_key", mongoKey),
		zap.String("instance_name", *wlResources.InstanceName),
		zap.Any("resource_data", resourceData))

	// 检查MongoDB连接
	if global.GVA_WANLIAN_MONGO == nil {
		global.GVA_LOG.Error("MongoDB连接为空")
		return errors.New("MongoDB连接未初始化")
	}

	// 插入MongoDB数据
	_, err = collection.InsertOne(ctx, mongoDoc)
	if err != nil {
		global.GVA_LOG.Error("MongoDB插入失败", zap.Error(err))
		tx.Rollback()
		return errors.Wrap(err, "MongoDB插入失败")
	}

	global.GVA_LOG.Info("MongoDB数据插入成功", zap.String("mongo_key", mongoKey))

	// 插入MySQL记录
	err = tx.Create(wlResources).Error
	if err != nil {
		global.GVA_LOG.Error("MySQL插入失败", zap.Error(err))
		tx.Rollback()
		return errors.Wrap(err, "MySQL插入失败")
	}

	global.GVA_LOG.Info("MySQL数据插入成功", zap.Uint("id", wlResources.ID))

	// 提交事务
	return tx.Commit().Error
}

// GetWlResourcesWithMongoData 获取资源信息（包含MongoDB数据）
func (wlResourcesService *WlResourcesService) GetWlResourcesWithMongoData(ctx context.Context, id uint) (wlResources *wl_playform.WlResources, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wlResources).Error
	if err != nil {
		return nil, err
	}

	// 从MongoDB获取详细数据
	if wlResources.ResourcesKey != nil {
		collection := global.GVA_WANLIAN_MONGO.Database.Collection("wl_resources")
		var mongoDoc bson.M
		err = collection.Find(ctx, bson.M{"mongo_key": *wlResources.ResourcesKey}).One(&mongoDoc)
		if err != nil {
			// MongoDB查询失败，但MySQL数据存在，返回MySQL数据
			global.GVA_LOG.Warn("MongoDB查询失败，返回MySQL数据", zap.Error(err))
			return wlResources, nil
		}
		// 这里可以将MongoDB数据合并到返回结果中
	}

	return wlResources, err
}

// UpdateWlResourcesWithTransaction 更新资源（事务处理）
func (wlResourcesService *WlResourcesService) UpdateWlResourcesWithTransaction(ctx context.Context, wlResources *wl_playform.WlResources, resourceData map[string]interface{}) error {
	// 开始MySQL事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新MongoDB数据
	if wlResources.ResourcesKey != nil {
		collection := global.GVA_WANLIAN_MONGO.Database.Collection("wl_resources")
		updateDoc := bson.M{
			"$set": bson.M{
				"instance_name": *wlResources.InstanceName,
				"resource_data": resourceData,
				"updated_at":    time.Now(),
			},
		}
		err := collection.UpdateOne(ctx, bson.M{"mongo_key": *wlResources.ResourcesKey}, updateDoc)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "MongoDB更新失败")
		}
	}

	// 更新MySQL记录
	err := tx.Save(wlResources).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "MySQL更新失败")
	}

	// 提交事务
	return tx.Commit().Error
}

// DeleteWlResourcesWithTransaction 删除资源（事务处理）
func (wlResourcesService *WlResourcesService) DeleteWlResourcesWithTransaction(ctx context.Context, wlResources *wl_playform.WlResources) error {
	// 开始MySQL事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除MongoDB数据
	if wlResources.ResourcesKey != nil {
		collection := global.GVA_WANLIAN_MONGO.Database.Collection("wl_resources")
		err := collection.Remove(ctx, bson.M{"mongo_key": *wlResources.ResourcesKey})
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "MongoDB删除失败")
		}
	}

	// 删除MySQL记录
	err := tx.Delete(wlResources).Error
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "MySQL删除失败")
	}

	// 提交事务
	return tx.Commit().Error
}

// CreateWlResources 创建资源
func (wlResourcesService *WlResourcesService) CreateWlResources(wlResources *wl_playform.WlResources) (err error) {
	err = global.GVA_DB.Create(wlResources).Error
	return err
}

// DeleteWlResources 删除资源信息
func (wlResourcesService *WlResourcesService) DeleteWlResources(wlResources *wl_playform.WlResources) (err error) {
	err = global.GVA_DB.Delete(wlResources).Error
	return err
}

// DeleteWlResourcesByIds 批量删除资源信息
func (wlResourcesService *WlResourcesService) DeleteWlResourcesByIds(IDs []uint) (err error) {
	err = global.GVA_DB.Delete(&[]wl_playform.WlResources{}, "id in ?", IDs).Error
	return err
}

// UpdateWlResources 更新资源信息
func (wlResourcesService *WlResourcesService) UpdateWlResources(wlResources *wl_playform.WlResources) (err error) {
	err = global.GVA_DB.Save(wlResources).Error
	return err
}

// GetWlResources 根据id获取资源信息
func (wlResourcesService *WlResourcesService) GetWlResources(id uint) (wlResources *wl_playform.WlResources, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&wlResources).Error
	return
}

// GetWlResourcesInfoList 分页获取资源信息
func (wlResourcesService *WlResourcesService) GetWlResourcesInfoList(info request.WlResourcesSearch) (list []*wl_playform.WlResources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&wl_playform.WlResources{})
	var wlResourcess []*wl_playform.WlResources
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.InstanceName != nil {
		db = db.Where("instance_name LIKE ?", "%"+*info.InstanceName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&wlResourcess).Error
	return wlResourcess, total, err
}

// VerifyWlResources 验证资源
func (wlResourcesService *WlResourcesService) VerifyWlResources(wlResources *wl_playform.WlResources) (err error) {
	// 模拟验证过程
	// 这里可以根据不同的资源类型进行不同的验证逻辑

	// 更新验证状态为成功
	status := "验证成功"
	wlResources.VerificationStatus = &status

	// 保存到数据库
	err = global.GVA_DB.Save(wlResources).Error
	if err != nil {
		return errors.Wrap(err, "更新验证状态失败")
	}

	return nil
}
