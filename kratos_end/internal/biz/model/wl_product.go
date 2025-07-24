package model

import "time"

type WlProducts struct {
	Id              int32     `gorm:"column:id;type:int;comment:主键id;primaryKey;not null;" json:"id"`                            // 主键id
	ProductsId      string    `gorm:"column:products_id;type:varchar(20);comment:产品id;default:NULL;" json:"products_id"`         // 产品id
	PrName          string    `gorm:"column:pr_name;type:varchar(20);comment:产品名称;default:NULL;" json:"pr_name"`                 // 产品名称
	PrCategory      string    `gorm:"column:pr_category;type:varchar(20);comment:所属品类;default:NULL;" json:"pr_category"`         // 所属品类
	StandardQuality int16     `gorm:"column:standard_quality;type:smallint;comment:标准品类;default:NULL;" json:"standard_quality"`  // 标准品类
	NodeType        string    `gorm:"column:node_type;type:varchar(20);comment:节点类型;default:NULL;" json:"node_type"`             // 节点类型
	AccessProtocol  string    `gorm:"column:access_protocol;type:varchar(20);comment:接入协议;default:NULL;" json:"access_protocol"` // 接入协议
	DataFormat      string    `gorm:"column:data_format;type:varchar(20);comment:数据格式;default:NULL;" json:"data_format"`         // 数据格式
	NetworkType     string    `gorm:"column:network_type;type:varchar(20);comment:网络类型;default:NULL;" json:"network_type"`       // 网络类型
	Factory         string    `gorm:"column:factory;type:varchar(100);comment:工厂;default:NULL;" json:"factory"`                  // 工厂
	PrInfo          string    `gorm:"column:pr_info;type:varchar(200);comment:产品描述;default:NULL;" json:"pr_info"`                // 产品描述
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	CreatedBy       uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;default:NULL;" json:"created_by"` // 创建者
	UpdatedBy       uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;default:NULL;" json:"updated_by"` // 更新者
	DeletedBy       uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;default:NULL;" json:"deleted_by"` // 删除者
}

func (WlProducts) TableName() string {
	return "wl_products"
}
