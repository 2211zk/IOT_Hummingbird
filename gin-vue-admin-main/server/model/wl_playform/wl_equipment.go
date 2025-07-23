
// 自动生成模板WlEquipment
package wl_playform
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlEquipment表 结构体  WlEquipment
type WlEquipment struct {
    global.GVA_MODEL
  EqName  *string `json:"eqName" form:"eqName" gorm:"comment:设备名称;column:eq_name;size:20;" binding:"required"`  //设备名称
  EqLogotype  *string `json:"eqLogotype" form:"eqLogotype" gorm:"comment:设备唯一标识;column:eq_logotype;size:50;"`  //设备唯一标识
  ProductsId  *int `json:"productsId" form:"productsId" gorm:"comment:所属产品;column:products_id;size:10;" binding:"required"`  //所属产品
  DriveId  *int `json:"driveId" form:"driveId" gorm:"comment:驱动id;column:drive_id;size:10;"`  //驱动id
  EqCoordinate  *string `json:"eqCoordinate" form:"eqCoordinate" gorm:"comment:设备坐标;column:eq_coordinate;size:50;"`  //设备坐标
  EqAddress  *string `json:"eqAddress" form:"eqAddress" gorm:"comment:设备详细地址;column:eq_address;size:100;"`  //设备详细地址
  EqInfo  *string `json:"eqInfo" form:"eqInfo" gorm:"comment:设备描述;column:eq_info;size:150;"`  //设备描述
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName wlEquipment表 WlEquipment自定义表名 wl_equipment
func (WlEquipment) TableName() string {
    return "wl_equipment"
}





