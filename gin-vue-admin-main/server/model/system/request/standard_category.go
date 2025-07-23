package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// StandardCategorySearch 标准品类搜索结构体
type StandardCategorySearch struct {
	request.PageInfo
	Name     string `json:"name" form:"name"`         // 品类名称
	Code     string `json:"code" form:"code"`         // 品类编码
	Category string `json:"category" form:"category"` // 所属类别
	Status   *int   `json:"status" form:"status"`     // 状态
}