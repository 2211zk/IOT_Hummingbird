package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// 产品表
type WlProducts struct {
	Id              int32     `gorm:"column:id;type:int;comment:主键id;primaryKey;" json:"id"`                          // 主键id
	ProductsId      string    `gorm:"column:products_id;type:varchar(20);comment:产品id;" json:"products_id"`           // 产品id
	PrName          string    `gorm:"column:pr_name;type:varchar(20);comment:产品名称;" json:"pr_name"`                 // 产品名称
	PrCategory      string    `gorm:"column:pr_category;type:varchar(20);comment:所属品类;" json:"pr_category"`         // 所属品类
	StandardQuality int16     `gorm:"column:standard_quality;type:smallint;comment:标准品类;" json:"standard_quality"`  // 标准品类
	NodeType        string    `gorm:"column:node_type;type:varchar(20);comment:节点类型;" json:"node_type"`             // 节点类型
	AccessProtocol  string    `gorm:"column:access_protocol;type:varchar(20);comment:接入协议;" json:"access_protocol"` // 接入协议
	DataFormat      string    `gorm:"column:data_format;type:varchar(20);comment:数据格式;" json:"data_format"`         // 数据格式
	NetworkType     string    `gorm:"column:network_type;type:varchar(20);comment:网络类型;" json:"network_type"`       // 网络类型
	Factory         string    `gorm:"column:factory;type:varchar(100);comment:工厂;" json:"factory"`                    // 工厂
	PrInfo          string    `gorm:"column:pr_info;type:varchar(200);comment:产品描述;" json:"pr_info"`                // 产品描述
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime(3);" json:"updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:datetime(3);" json:"deleted_at"`
	CreatedBy       uint64    `gorm:"column:created_by;type:bigint UNSIGNED;comment:创建者;" json:"created_by"` // 创建者
	UpdatedBy       uint64    `gorm:"column:updated_by;type:bigint UNSIGNED;comment:更新者;" json:"updated_by"` // 更新者
	DeletedBy       uint64    `gorm:"column:deleted_by;type:bigint UNSIGNED;comment:删除者;" json:"deleted_by"` // 删除者
}

// 定义结构体名字
func (w *WlProducts) TableName() string {
	return "wl_products"
}

type WlProductsRepo interface {
	Save(context.Context, *WlProducts) (*WlProducts, error)
	ListAll(context.Context) ([]*WlProducts, error)
}

type WlProductsService struct {
	repo WlProductsRepo
	log  *log.Helper
}

func NewWlProductsUsecase(repo WlProductsRepo, logger log.Logger) *WlProductsService {
	return &WlProductsService{repo: repo, log: log.NewHelper(logger)}
}

func (s *WlProductsService) Save(ctx context.Context, p *WlProducts) (*WlProducts, error) {
	return s.repo.Save(ctx, p)
}

func (s *WlProductsService) ListAll(ctx context.Context) ([]*WlProducts, error) {
	return s.repo.ListAll(ctx)
}
