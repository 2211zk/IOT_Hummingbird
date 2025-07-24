package model

import "time"

type WlAlarm struct {
	Id           int64     `gorm:"column:id;type:bigint;comment:告警ID;primaryKey;not null;" json:"id"`                                       // 告警ID
	DeviceId     int64     `gorm:"column:device_id;type:bigint;comment:设备ID;not null;" json:"device_id"`                                    // 设备ID
	AlarmType    string    `gorm:"column:alarm_type;type:varchar(50);comment:告警类型（如：设备离线、数据异常、阈值超限等）;not null;" json:"alarm_type"`          // 告警类型（如：设备离线、数据异常、阈值超限等）
	AlarmLevel   int8      `gorm:"column:alarm_level;type:tinyint;comment:告警级别：1-低，2-中，3-高，4-紧急;not null;default:1;" json:"alarm_level"`    // 告警级别：1-低，2-中，3-高，4-紧急
	AlarmStatus  int8      `gorm:"column:alarm_status;type:tinyint;comment:告警状态：0-未处理，1-已处理，2-已忽略;not null;default:0;" json:"alarm_status"` // 告警状态：0-未处理，1-已处理，2-已忽略
	AlarmContent string    `gorm:"column:alarm_content;type:text;comment:告警内容描述;" json:"alarm_content"`                                     // 告警内容描述
	AlarmData    string    `gorm:"column:alarm_data;type:json;comment:告警相关数据（JSON格式）;default:NULL;" json:"alarm_data"`                      // 告警相关数据（JSON格式）
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;comment:告警创建时间;not null;default:CURRENT_TIMESTAMP;" json:"create_time"`  // 告警创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;comment:告警更新时间;not null;default:CURRENT_TIMESTAMP;" json:"update_time"`  // 告警更新时间
	HandleTime   time.Time `gorm:"column:handle_time;type:datetime;comment:处理时间;default:NULL;" json:"handle_time"`                          // 处理时间
	HandleUser   string    `gorm:"column:handle_user;type:varchar(50);comment:处理人;default:NULL;" json:"handle_user"`                        // 处理人
	HandleRemark string    `gorm:"column:handle_remark;type:text;comment:处理备注;" json:"handle_remark"`                                       // 处理备注
}

func (WlAlarm) TableName() string {
	return "wl_alarm"
}
