// 自动生成模板WlEngineRules
package wl_playform

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// wlEngineRules表 结构体  WlEngineRules
type WlEngineRules struct {
	global.GVA_MODEL
	RuleName         *string `json:"ruleName" form:"ruleName" gorm:"comment:规则名称;column:rule_name;size:30;" binding:"required"`                         //规则名称
	RuleDescription  *string `json:"ruleDescription" form:"ruleDescription" gorm:"comment:规则描述;column:rule_description;size:200;"`                      //规则描述
	MessageSource    *string `json:"messageSource" form:"messageSource" gorm:"comment:消息源;column:message_source;size:100;" binding:"required"`          //消息源
	QueryField       *string `json:"queryField" form:"queryField" gorm:"comment:查询字段;column:query_field;size:100;" binding:"required"`                  //查询字段
	Condition        *string `json:"condition" form:"condition" gorm:"comment:条件;column:condition;size:100;"`                                           //条件
	SqlStatement     *string `json:"sqlStatement" form:"sqlStatement" gorm:"comment:sql语句;column:sql_statement;size:255;"`                              //sql语句
	ForwardingMethod *string `json:"forwardingMethod" form:"forwardingMethod" gorm:"comment:转换方法;column:forwarding_method;size:20;" binding:"required"` //转换方法
	ResourceId       *int    `json:"resourceId" form:"resourceId" gorm:"comment:使用资源id;column:resource_id;size:10;" binding:"required"`                 //使用资源id
	RuleStatus       *string `json:"ruleStatus" form:"ruleStatus" gorm:"comment:启用状态;column:rule_status;size:20;default:'0'"`                           //启用状态
	CreatedBy        uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName wlEngineRules表 WlEngineRules自定义表名 wl_engine_rules
func (WlEngineRules) TableName() string {
	return "wl_engine_rules"
}
