
// 自动生成模板WlProducts
package wl_playform
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlProducts表 结构体  WlProducts
type WlProducts struct {
    global.GVA_MODEL
  PrName  *string `json:"prName" form:"prName" gorm:"comment:产品名称;column:pr_name;size:20;" binding:"required"`  //产品名称
  PrCategory  *string `json:"prCategory" form:"prCategory" gorm:"comment:所属品类;column:pr_category;size:20;" binding:"required"`  //所属品类
  StandardQuality  *int `json:"standardQuality" form:"standardQuality" gorm:"comment:标准品类;column:standard_quality;size:10;" binding:"required"`  //标准品类
  NodeType  *string `json:"nodeType" form:"nodeType" gorm:"comment:节点类型;column:node_type;size:20;" binding:"required"`  //节点类型
  AccessProtocol  *string `json:"accessProtocol" form:"accessProtocol" gorm:"comment:接入协议;column:access_protocol;size:20;" binding:"required"`  //接入协议
  DataFormat  *string `json:"dataFormat" form:"dataFormat" gorm:"comment:数据格式;column:data_format;size:20;" binding:"required"`  //数据格式
  NetworkType  *string `json:"networkType" form:"networkType" gorm:"comment:网络类型;column:network_type;size:20;" binding:"required"`  //网络类型
  Factory  *string `json:"factory" form:"factory" gorm:"comment:工厂;column:factory;size:100;"`  //工厂
  PrInfo  *string `json:"prInfo" form:"prInfo" gorm:"comment:产品描述;column:pr_info;size:200;"`  //产品描述
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wlProducts表 WlProducts自定义表名 wl_products
func (WlProducts) TableName() string {
    return "wl_products"
}





