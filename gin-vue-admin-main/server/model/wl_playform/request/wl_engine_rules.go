
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WlEngineRulesSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      RuleName  *string `json:"ruleName" form:"ruleName"` 
      MessageSource  *string `json:"messageSource" form:"messageSource"` 
    request.PageInfo
}
