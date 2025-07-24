package model

import "time"

type WlDeviceMessage struct {
	Id               int64     `gorm:"column:id;type:bigint;comment:消息ID;primaryKey;not null;" json:"id"`                                                       // 消息ID
	DeviceId         int64     `gorm:"column:device_id;type:bigint;comment:设备ID;not null;" json:"device_id"`                                                    // 设备ID
	MessageType      string    `gorm:"column:message_type;type:varchar(50);comment:消息类型（如：数据上报、状态变更、命令响应等）;not null;" json:"message_type"`                      // 消息类型（如：数据上报、状态变更、命令响应等）
	MessageDirection int8      `gorm:"column:message_direction;type:tinyint;comment:消息方向：1-上行（设备到平台），2-下行（平台到设备）;not null;default:1;" json:"message_direction"` // 消息方向：1-上行（设备到平台），2-下行（平台到设备）
	MessageContent   string    `gorm:"column:message_content;type:text;comment:消息内容;" json:"message_content"`                                                   // 消息内容
	MessageData      string    `gorm:"column:message_data;type:json;comment:消息数据（JSON格式）;default:NULL;" json:"message_data"`                                    // 消息数据（JSON格式）
	MessageSize      int32     `gorm:"column:message_size;type:int;comment:消息大小（字节）;default:0;" json:"message_size"`                                            // 消息大小（字节）
	Status           int8      `gorm:"column:status;type:tinyint;comment:消息状态：1-成功，2-失败，3-超时;not null;default:1;" json:"status"`                                // 消息状态：1-成功，2-失败，3-超时
	CreateTime       time.Time `gorm:"column:create_time;type:datetime;comment:消息创建时间;not null;default:CURRENT_TIMESTAMP;" json:"create_time"`                  // 消息创建时间
	UpdateTime       time.Time `gorm:"column:update_time;type:datetime;comment:消息更新时间;not null;default:CURRENT_TIMESTAMP;" json:"update_time"`                  // 消息更新时间
}

func (WlDeviceMessage) TableName() string {
	return "wl_device_message"
}
