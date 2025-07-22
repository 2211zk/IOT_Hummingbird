
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WlEquipmentSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      EqName  *string `json:"eqName" form:"eqName"` 
      ProductsId  *int `json:"productsId" form:"productsId"` 
    request.PageInfo
}
