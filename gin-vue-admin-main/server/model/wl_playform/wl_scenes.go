
// 自动生成模板WlScenes
package wl_playform
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlScenes表 结构体  WlScenes
type WlScenes struct {
    global.GVA_MODEL
  SceneName  *string `json:"sceneName" form:"sceneName" gorm:"comment:场景名称;column:scene_name;size:20;" binding:"required"`  //场景名称
  ScenesDescription  *string `json:"scenesDescription" form:"scenesDescription" gorm:"comment:场景描述;column:scenes_description;size:100;"`  //场景描述
  ScenesStatus  *string `json:"scenesStatus" form:"scenesStatus" gorm:"comment:启动状态;column:scenes_status;size:20;"`  //启动状态
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wlScenes表 WlScenes自定义表名 wl_scenes
func (WlScenes) TableName() string {
    return "wl_scenes"
}





