package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WlDriversSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	DriverName     string      `json:"driverName" form:"driverName"`
	DriverType     string      `json:"driverType" form:"driverType"`
	request.PageInfo
}
