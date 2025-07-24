package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_driver"
)

type WlDriverMarketSearch struct {
	wl_driver.WlDriverMarket
	request.PageInfo
	CreatedAtRange []string `json:"createdAtRange[]" form:"createdAtRange[]"`
}
