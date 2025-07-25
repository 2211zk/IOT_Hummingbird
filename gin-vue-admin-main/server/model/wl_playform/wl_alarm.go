// 自动生成模板WlAlarm
package wl_playform

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlAlarm表 结构体  WlAlarm
type WlAlarm struct {
	global.GVA_MODEL
	DeviceId     *int       `json:"deviceId" form:"deviceId" gorm:"comment:设备ID;column:device_id;" binding:"required"`     //设备ID
	AlarmType    *string    `json:"alarmType" form:"alarmType" gorm:"comment:告警类型;column:alarm_type;size:50;"`             //告警类型
	AlarmLevel   *string    `json:"alarmLevel" form:"alarmLevel" gorm:"comment:告警级别;column:alarm_level;size:20;"`          //告警级别
	AlarmStatus  *string    `json:"alarmStatus" form:"alarmStatus" gorm:"comment:告警状态;column:alarm_status;size:20;"`       //告警状态
	AlarmContent *string    `json:"alarmContent" form:"alarmContent" gorm:"comment:告警内容描述;column:alarm_content;size:191;"` //告警内容描述
	AlarmData    *string    `json:"alarmData" form:"alarmData" gorm:"comment:告警相关数据;column:alarm_data;size:191;"`          //告警相关数据
	CreateTime   *time.Time `json:"createTime" form:"createTime" gorm:"comment:告警创建时间;column:create_time;"`                //告警创建时间
	UpdateTime   *time.Time `json:"updateTime" form:"updateTime" gorm:"comment:告警更新时间;column:update_time;"`                //告警更新时间
	HandleTime   *time.Time `json:"handleTime" form:"handleTime" gorm:"comment:处理时间;column:handle_time;"`                  //处理时间
	HandleUser   *string    `json:"handleUser" form:"handleUser" gorm:"comment:处理人;column:handle_user;size:50;"`           //处理人
	HandleRemark *string    `json:"handleRemark" form:"handleRemark" gorm:"comment:处理备注;column:handle_remark;size:191;"`   //处理备注
	CreatedBy    *int       `json:"createdBy" form:"createdBy" gorm:"comment:创建者;column:created_by;size:20;"`              //创建者
	UpdatedBy    *int       `json:"updatedBy" form:"updatedBy" gorm:"comment:更新者;column:updated_by;size:20;"`              //更新者
	DeletedBy    *int       `json:"deletedBy" form:"deletedBy" gorm:"comment:删除者;column:deleted_by;size:20;"`              //删除者
}

// TableName wlAlarm表 WlAlarm自定义表名 wl_alarm
func (WlAlarm) TableName() string {
	return "wl_alarm"
}
