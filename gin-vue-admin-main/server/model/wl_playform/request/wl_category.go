
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WlCategorySearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    Name           *string     `json:"name" form:"name"`           // 品类名称
    CaName         *string     `json:"caName" form:"caName"`       // 品类名称（后端字段）
    request.PageInfo
}
