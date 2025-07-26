package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DriverCardsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           string      `json:"name" form:"name"`
	Tags           string      `json:"tags" form:"tags"`
	Description    string      `json:"description" form:"description"`
	request.PageInfo
}
