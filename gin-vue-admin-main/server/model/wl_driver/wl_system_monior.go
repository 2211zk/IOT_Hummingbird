package wl_driver

import (
	"time"
)

type WlSystemMonior struct {
	ID             uint64     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Timestamp      *time.Time `json:"timestamp" gorm:"column:timestamp"`
	CpuUsage       *float64   `json:"cpu_usage" gorm:"column:cpu_usage;type:decimal(10,2);comment:CPU使用率"`
	MemoryUpRate   *int       `json:"memory_up_rate" gorm:"column:memory_up_rate;comment:上行消息速率"`
	MemoryDownRate *int       `json:"memory_down_rate" gorm:"column:memory_down_rate;comment:下行消息速率"`
}

func (WlSystemMonior) TableName() string {
	return "wl_system_monior"
}
