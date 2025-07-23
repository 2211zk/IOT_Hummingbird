
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WlScenesSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      SceneName  *string `json:"sceneName" form:"sceneName"` 
      ScenesStatus  *string `json:"scenesStatus" form:"scenesStatus"` 
    request.PageInfo
}
