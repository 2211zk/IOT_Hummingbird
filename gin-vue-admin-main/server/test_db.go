package main

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

func main() {
	// 初始化配置
	initialize.Config()

	// 初始化数据库
	initialize.DB()

	// 测试数据库连接
	var count int64
	err := global.GVA_DB.Model(&struct{}{}).Raw("SELECT 1").Count(&count).Error
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	fmt.Println("数据库连接成功!")

	// 检查wl_department表是否存在
	var tableExists bool
	err = global.GVA_DB.Raw("SELECT COUNT(*) > 0 FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'wl_department'").Scan(&tableExists).Error
	if err != nil {
		log.Printf("检查表失败: %v", err)
	} else if tableExists {
		fmt.Println("wl_department表存在!")

		// 检查表中的数据
		var count int64
		err = global.GVA_DB.Model(&struct{}{}).Raw("SELECT COUNT(*) FROM wl_department").Count(&count).Error
		if err != nil {
			log.Printf("查询数据失败: %v", err)
		} else {
			fmt.Printf("wl_department表中有 %d 条数据\n", count)
		}
	} else {
		fmt.Println("wl_department表不存在!")
	}
}
