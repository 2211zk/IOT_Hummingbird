package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WlEquipmentSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	EqName         *string     `json:"eqName" form:"eqName"`
	ProductsId     int         `json:"productsId" form:"productsId"`
	Status         *string     `json:"status" form:"status"`
	request.PageInfo
}
