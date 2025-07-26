
// 自动生成模板WlUser
package wl_playform
import (
	"time"
)

// wlUser表 结构体  WlUser
type WlUser struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;comment:主键id;column:id;size:10;"`  //主键id
  UserName  *string `json:"userName" form:"userName" gorm:"comment:用户名;column:user_name;size:20;"`  //用户名
  UserNickname  *string `json:"userNickname" form:"userNickname" gorm:"comment:用户昵称;column:user_nickname;size:20;"`  //用户昵称
  Department  *int `json:"department" form:"department" gorm:"comment:部门;column:department;size:10;"`  //部门
  Mobile  *string `json:"mobile" form:"mobile" gorm:"comment:手机号;column:mobile;"`  //手机号
  Email  *string `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:100;"`  //邮箱
  Password  *string `json:"password" form:"password" gorm:"comment:密码;column:password;size:100;"`  //密码
  Gender  *string `json:"gender" form:"gender" gorm:"comment:性别;column:gender;size:20;"`  //性别
  Role  *int `json:"role" form:"role" gorm:"comment:角色;column:role;size:10;"`  //角色
  UserStatus  *string `json:"userStatus" form:"userStatus" gorm:"comment:状态;column:user_status;size:20;"`  //状态
  Comment  *string `json:"comment" form:"comment" gorm:"comment:备注;column:comment;size:200;"`  //备注
  CreationTime  *time.Time `json:"creationTime" form:"creationTime" gorm:"comment:创建时间;column:creation_time;"`  //创建时间
}


// TableName wlUser表 WlUser自定义表名 wl_user
func (WlUser) TableName() string {
    return "wl_user"
}





