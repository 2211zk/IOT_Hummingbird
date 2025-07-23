
// 自动生成模板WlCaFunction
package wl_playform
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlCaFunction表 结构体  WlCaFunction
type WlCaFunction struct {
    global.GVA_MODEL
  CaId  *int `json:"caId" form:"caId" gorm:"comment:品类id;column:ca_id;size:10;"`  //品类id
  FunctionType  *string `json:"functionType" form:"functionType" gorm:"comment:功能类型;column:function_type;size:20;"`  //功能类型
  FunctionName  *string `json:"functionName" form:"functionName" gorm:"comment:功能名称;column:function_name;size:20;"`  //功能名称
  Identifier  *string `json:"identifier" form:"identifier" gorm:"comment:标识符;column:identifier;size:50;"`  //标识符
  DataType  *string `json:"dataType" form:"dataType" gorm:"comment:数据类型;column:data_type;size:20;"`  //数据类型
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wlCaFunction表 WlCaFunction自定义表名 wl_ca_function
func (WlCaFunction) TableName() string {
    return "wl_ca_function"
}





