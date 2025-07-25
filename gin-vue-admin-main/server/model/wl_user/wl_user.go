package wl_user

import "time"

type WlUser struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	user_name     string    `json:"userName" gorm:"column:user_name"`
	user_nickname string    `json:"userNickname" gorm:"column:user_nickname"`
	Department    int       `json:"department"`
	Mobile        string    `json:"mobile"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Gender        string    `json:"gender"`
	Role          int       `json:"role"`
	UserStatus    string    `json:"userStatus" gorm:"column:user_status"`
	Comment       string    `json:"comment"`
	CreationTime  time.Time `json:"creationTime" gorm:"column:creation_time"`
}

func (WlUser) TableName() string {
	return "wl_user"
}
