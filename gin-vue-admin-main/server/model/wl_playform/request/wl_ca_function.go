
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WlCaFunctionSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    CaId           *int        `json:"caId" form:"caId"`           // 品类ID
    FunctionName   *string     `json:"functionName" form:"functionName"`   // 功能名称
    FunctionType   *string     `json:"functionType" form:"functionType"`   // 功能类型
    request.PageInfo
}
