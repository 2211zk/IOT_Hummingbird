// 自动生成模板WlDrivers
package wl_driver

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// wlDrivers表 结构体  WlDrivers
type WlDrivers struct {
	global.GVA_MODEL
	DriverNum      *string    `json:"driverNum" form:"driverNum" gorm:"comment:驱动编号;column:driver_num;size:50;"`                //驱动编号
	DriverName     *string    `json:"driverName" form:"driverName" gorm:"comment:驱动名称;column:driver_name;size:100;"`            //驱动名称
	Version        *string    `json:"version" form:"version" gorm:"comment:版本;column:version;size:20;"`                         //版本
	DriverType     *string    `json:"driverType" form:"driverType" gorm:"comment:驱动类型;column:driver_type;size:50;"`             //驱动类型
	Status         *string    `json:"status" form:"status" gorm:"comment:状态;column:status;size:1;"`                             //状态
	CreatedTime    *time.Time `json:"createdTime" form:"createdTime" gorm:"comment:创建时间;column:created_time;"`                  //创建时间
	UpdateTime     *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:更新时间;column:update_time;"`                     //更新时间
	ProtocolType   *string    `json:"protocolType" form:"protocolType" gorm:"comment:协议类型;column:protocol_type;size:50;"`       //协议类型
	DeviceCategory *string    `json:"deviceCategory" form:"deviceCategory" gorm:"comment:设备类型;column:device_category;size:50;"` //设备类型
	DriverId       *string    `json:"driverId" form:"driverId" gorm:"comment:驱动编号;column:driver_id;size:20;"`                   //驱动编号
	CreatedBy      *int       `json:"createdBy" form:"createdBy" gorm:"comment:创建者;column:created_by;size:20;"`                 //创建者
	UpdatedBy      *int       `json:"updatedBy" form:"updatedBy" gorm:"comment:更新者;column:updated_by;size:20;"`                 //更新者
	DeletedBy      *int       `json:"deletedBy" form:"deletedBy" gorm:"comment:删除者;column:deleted_by;size:20;"`                 //删除者
}

// TableName wlDrivers表 WlDrivers自定义表名 wl_drivers
func (WlDrivers) TableName() string {
	return "wl_drivers"
}
