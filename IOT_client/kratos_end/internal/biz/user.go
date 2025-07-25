package biz

type User struct {
	Id           int32
	UserName     string
	UserNickname string
	Department   int32
	Mobile       string
	Email        string
	Password     string
	Gender       string
	Role         int32
	UserStatus   string
	Comment      string
}

// 用户注册业务接口（支持DTM分布式事务）
type UserRepo interface {
	Register(user *User, dtmGid string) (int32, error)
	RegisterCompensate(id int32, dtmGid string) error
	// 其它常规接口
	Login(userName, password string) (*User, error)
	GetUser(id int32) (*User, error)
	ListUser(page, pageSize int32) ([]*User, int32, error)
	UpdateUser(user *User) error
	DeleteUser(id int32) error
}
