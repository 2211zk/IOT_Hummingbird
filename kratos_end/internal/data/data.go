package data

import (
	"kratos/internal/conf"

	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"database/sql"
	"kratos/internal/biz"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	MongoClient *mongo.Client
	MySQLClient *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 构建MongoDB连接URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		c.Mongodb.User,
		c.Mongodb.Password,
		c.Mongodb.Host,
		c.Mongodb.Port,
		c.Mongodb.Database,
	)
	clientOpts := options.Client().ApplyURI(uri)
	mongoClient, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, nil, err
	}
	// 初始化MySQL
	mysqlDB, err := sql.Open("mysql", c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = mongoClient.Disconnect(context.Background())
		_ = mysqlDB.Close()
	}
	return &Data{MongoClient: mongoClient, MySQLClient: mysqlDB}, cleanup, nil
}

// UserRepo实现
type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Register(user *biz.User, dtmGid string) (int32, error) {
	res, err := r.data.MySQLClient.Exec(`INSERT INTO wl_user (user_name, user_nickname, department, mobile, email, password, gender, role, user_status, comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.UserName, user.UserNickname, user.Department, user.Mobile, user.Email, user.Password, user.Gender, user.Role, user.UserStatus, user.Comment)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int32(id), err
}
func (r *userRepo) RegisterCompensate(id int32, dtmGid string) error {
	_, err := r.data.MySQLClient.Exec(`DELETE FROM wl_user WHERE id = ?`, id)
	return err
}
func (r *userRepo) Login(userName, password string) (*biz.User, error) {
	row := r.data.MySQLClient.QueryRow(`SELECT id, user_name, user_nickname, department, mobile, email, password, gender, role, user_status, comment FROM wl_user WHERE user_name = ? AND password = ?`, userName, password)
	var user biz.User
	if err := row.Scan(&user.Id, &user.UserName, &user.UserNickname, &user.Department, &user.Mobile, &user.Email, &user.Password, &user.Gender, &user.Role, &user.UserStatus, &user.Comment); err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) GetUser(id int32) (*biz.User, error) {
	row := r.data.MySQLClient.QueryRow(`SELECT id, user_name, user_nickname, department, mobile, email, password, gender, role, user_status, comment FROM wl_user WHERE id = ?`, id)
	var user biz.User
	if err := row.Scan(&user.Id, &user.UserName, &user.UserNickname, &user.Department, &user.Mobile, &user.Email, &user.Password, &user.Gender, &user.Role, &user.UserStatus, &user.Comment); err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) ListUser(page, pageSize int32) ([]*biz.User, int32, error) {
	offset := (page - 1) * pageSize
	rows, err := r.data.MySQLClient.Query(`SELECT id, user_name, user_nickname, department, mobile, email, password, gender, role, user_status, comment FROM wl_user LIMIT ? OFFSET ?`, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var users []*biz.User
	for rows.Next() {
		var user biz.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.UserNickname, &user.Department, &user.Mobile, &user.Email, &user.Password, &user.Gender, &user.Role, &user.UserStatus, &user.Comment); err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}
	var total int32
	row := r.data.MySQLClient.QueryRow(`SELECT COUNT(*) FROM wl_user`)
	_ = row.Scan(&total)
	return users, total, nil
}
func (r *userRepo) UpdateUser(user *biz.User) error {
	_, err := r.data.MySQLClient.Exec(`UPDATE wl_user SET user_nickname=?, department=?, mobile=?, email=?, gender=?, role=?, user_status=?, comment=? WHERE id=?`,
		user.UserNickname, user.Department, user.Mobile, user.Email, user.Gender, user.Role, user.UserStatus, user.Comment, user.Id)
	return err
}
func (r *userRepo) DeleteUser(id int32) error {
	_, err := r.data.MySQLClient.Exec(`DELETE FROM wl_user WHERE id = ?`, id)
	return err
}
