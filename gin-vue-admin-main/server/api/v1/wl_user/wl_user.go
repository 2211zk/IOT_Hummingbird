package wl_user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wl_user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wl_user"
	"github.com/gin-gonic/gin"
)

var WlUserApi = new(WlUserApiStruct)

type WlUserApiStruct struct{}

func (a *WlUserApiStruct) GetWlUserList(c *gin.Context) {
	var req request.WlUserSearch
	_ = c.ShouldBindJSON(&req)
	list, total, err := wl_user.WlUserServiceApp.GetWlUserList(req)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "查询成功", c)
}
