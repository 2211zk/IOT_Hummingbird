
// 自动生成模板WlCategory
package wl_playform
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlCategory表 结构体  WlCategory
type WlCategory struct {
    global.GVA_MODEL
  CaName  *string `json:"caName,name" form:"caName" gorm:"comment:品类名称;column:ca_name;size:50;"`  //品类名称
  CaKey  *string `json:"caKey,key" form:"caKey" gorm:"comment:品类键;column:ca_key;size:20;"`  //品类键
  CaScenario  *string `json:"caScenario,scenario" form:"caScenario" gorm:"comment:所属场景;column:ca_scenario;size:20;"`  //所属场景
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wlCategory表 WlCategory自定义表名 wl_category
func (WlCategory) TableName() string {
    return "wl_category"
}





