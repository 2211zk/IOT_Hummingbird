package performance

import (
	"fmt"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_department/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_department"
	"github.com/stretchr/testify/assert"
)

// TestDepartmentListPerformance 测试部门列表查询性能
func TestDepartmentListPerformance(t *testing.T) {
	service := wl_department.WlDepartmentServiceApp

	// 测试树形模式性能
	t.Run("TreeModePerformance", func(t *testing.T) {
		req := request.WlDepartmentSearch{
			TreeMode: true,
			Page:     1,
			PageSize: 100,
		}

		start := time.Now()
		list, total, err := service.GetWlDepartmentList(req)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.GreaterOrEqual(t, total, int64(0))

		// 性能断言：树形查询应该在500ms内完成
		assert.Less(t, duration, 500*time.Millisecond,
			fmt.Sprintf("Tree mode query took %v, expected < 500ms", duration))

		t.Logf("Tree mode query completed in %v with %d records", duration, total)
	})

	// 测试平铺模式性能
	t.Run("FlatModePerformance", func(t *testing.T) {
		req := request.WlDepartmentSearch{
			TreeMode: false,
			Page:     1,
			PageSize: 20,
		}

		start := time.Now()
		list, total, err := service.GetWlDepartmentList(req)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.GreaterOrEqual(t, total, int64(0))

		// 性能断言：平铺查询应该在200ms内完成
		assert.Less(t, duration, 200*time.Millisecond,
			fmt.Sprintf("Flat mode query took %v, expected < 200ms", duration))

		t.Logf("Flat mode query completed in %v with %d records", duration, total)
	})

	// 测试搜索性能
	t.Run("SearchPerformance", func(t *testing.T) {
		req := request.WlDepartmentSearch{
			Name:     "测试",
			TreeMode: false,
			Page:     1,
			PageSize: 20,
		}

		start := time.Now()
		list, total, err := service.GetWlDepartmentList(req)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.GreaterOrEqual(t, total, int64(0))

		// 性能断言：搜索查询应该在300ms内完成
		assert.Less(t, duration, 300*time.Millisecond,
			fmt.Sprintf("Search query took %v, expected < 300ms", duration))

		t.Logf("Search query completed in %v with %d records", duration, total)
	})
}

// TestDepartmentTreePerformance 测试部门树查询性能
func TestDepartmentTreePerformance(t *testing.T) {
	service := wl_department.WlDepartmentServiceApp

	t.Run("DepartmentTreePerformance", func(t *testing.T) {
		req := request.DepartmentTreeReq{}

		start := time.Now()
		tree, err := service.GetDepartmentTree(req)
		duration := time.Since(start)

		assert.NoError(t, err)
		assert.NotNil(t, tree)

		// 性能断言：部门树查询应该在100ms内完成
		assert.Less(t, duration, 100*time.Millisecond,
			fmt.Sprintf("Department tree query took %v, expected < 100ms", duration))

		t.Logf("Department tree query completed in %v with %d nodes", duration, len(tree))
	})

	// 测试缓存性能
	t.Run("CachePerformance", func(t *testing.T) {
		req := request.DepartmentTreeReq{}

		// 第一次查询（无缓存）
		start1 := time.Now()
		tree1, err1 := service.GetDepartmentTree(req)
		duration1 := time.Since(start1)

		assert.NoError(t, err1)
		assert.NotNil(t, tree1)

		// 第二次查询（有缓存）
		start2 := time.Now()
		tree2, err2 := service.GetDepartmentTree(req)
		duration2 := time.Since(start2)

		assert.NoError(t, err2)
		assert.NotNil(t, tree2)

		// 缓存查询应该更快
		assert.Less(t, duration2, duration1,
			fmt.Sprintf("Cached query (%v) should be faster than non-cached (%v)", duration2, duration1))

		t.Logf("Non-cached query: %v, Cached query: %v, Improvement: %v",
			duration1, duration2, duration1-duration2)
	})
}

// TestConcurrentPerformance 测试并发性能
func TestConcurrentPerformance(t *testing.T) {
	service := wl_department.WlDepartmentServiceApp
	concurrency := 10
	iterations := 100

	t.Run("ConcurrentQueries", func(t *testing.T) {
		results := make(chan time.Duration, concurrency*iterations)
		errors := make(chan error, concurrency*iterations)

		start := time.Now()

		// 启动并发查询
		for i := 0; i < concurrency; i++ {
			go func() {
				for j := 0; j < iterations; j++ {
					queryStart := time.Now()

					req := request.WlDepartmentSearch{
						TreeMode: true,
						Page:     1,
						PageSize: 50,
					}

					_, _, err := service.GetWlDepartmentList(req)
					queryDuration := time.Since(queryStart)

					if err != nil {
						errors <- err
					} else {
						results <- queryDuration
					}
				}
			}()
		}

		// 收集结果
		var totalDuration time.Duration
		var maxDuration time.Duration
		var minDuration time.Duration = time.Hour
		successCount := 0
		errorCount := 0

		for i := 0; i < concurrency*iterations; i++ {
			select {
			case duration := <-results:
				totalDuration += duration
				if duration > maxDuration {
					maxDuration = duration
				}
				if duration < minDuration {
					minDuration = duration
				}
				successCount++
			case err := <-errors:
				t.Logf("Query error: %v", err)
				errorCount++
			case <-time.After(30 * time.Second):
				t.Fatal("Timeout waiting for concurrent queries")
			}
		}

		totalTime := time.Since(start)
		avgDuration := totalDuration / time.Duration(successCount)

		// 性能断言
		assert.Equal(t, 0, errorCount, "No errors should occur during concurrent queries")
		assert.Equal(t, concurrency*iterations, successCount, "All queries should succeed")
		assert.Less(t, avgDuration, 1*time.Second, "Average query time should be < 1s")
		assert.Less(t, maxDuration, 2*time.Second, "Max query time should be < 2s")

		t.Logf("Concurrent performance results:")
		t.Logf("  Total time: %v", totalTime)
		t.Logf("  Success rate: %d/%d", successCount, concurrency*iterations)
		t.Logf("  Average query time: %v", avgDuration)
		t.Logf("  Min query time: %v", minDuration)
		t.Logf("  Max query time: %v", maxDuration)
		t.Logf("  Queries per second: %.2f", float64(successCount)/totalTime.Seconds())
	})
}

// TestMemoryUsage 测试内存使用情况
func TestMemoryUsage(t *testing.T) {
	service := wl_department.WlDepartmentServiceApp

	t.Run("MemoryUsage", func(t *testing.T) {
		// 执行多次查询以观察内存使用
		for i := 0; i < 100; i++ {
			req := request.WlDepartmentSearch{
				TreeMode: true,
				Page:     1,
				PageSize: 100,
			}

			_, _, err := service.GetWlDepartmentList(req)
			assert.NoError(t, err)

			// 每10次查询检查一次内存
			if i%10 == 0 {
				// 这里可以添加内存使用检查逻辑
				// 例如使用 runtime.ReadMemStats
				t.Logf("Completed %d queries", i+1)
			}
		}

		t.Log("Memory usage test completed")
	})
}

// BenchmarkDepartmentList 基准测试
func BenchmarkDepartmentList(b *testing.B) {
	service := wl_department.WlDepartmentServiceApp

	b.Run("TreeMode", func(b *testing.B) {
		req := request.WlDepartmentSearch{
			TreeMode: true,
			Page:     1,
			PageSize: 100,
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := service.GetWlDepartmentList(req)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("FlatMode", func(b *testing.B) {
		req := request.WlDepartmentSearch{
			TreeMode: false,
			Page:     1,
			PageSize: 20,
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _, err := service.GetWlDepartmentList(req)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkDepartmentTree 部门树基准测试
func BenchmarkDepartmentTree(b *testing.B) {
	service := wl_department.WlDepartmentServiceApp
	req := request.DepartmentTreeReq{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.GetDepartmentTree(req)
		if err != nil {
			b.Fatal(err)
		}
	}
}
