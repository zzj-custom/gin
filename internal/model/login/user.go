package login

import (
	"go-api/internal/model"
	"go-api/internal/util"
	"go-api/pkg/mysql"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type User struct {
	model.Model
	UserName  string `gorm:"type:varchar(20);comment:用户名;not null;index:idx_user_name;unique"`
	Password  string `gorm:"type:char(32);comment:密码;not null"`
	SecretKey string `gorm:"type:varchar(64);comment:秘钥;not null"`
	AppId     int64  `gorm:"type:bigint(32);comment:app_id;not null"`
	Email     string `gorm:"type:varchar(20);comment:邮箱;"`
	Phone     int64  `gorm:"type:bigint(11);comment:手机号;"`
	Ip        string `gorm:"type:varchar(20);comment:ip地址;"`
	Status    uint8  `gorm:"type:tinyint(3);comment:状态;default:1;not null"`
}

var (
	userRepo     *userRepository
	userRepoOnce sync.Once
)

func (u *User) TableName() string {
	return "user"
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo() *userRepository {
	userRepoOnce.Do(func() {
		userRepo = new(userRepository)
		db, err := mysql.GetConnect("go")
		if err != nil {
			return
		}
		userRepo.db = db
	})
	return userRepo
}

func (ur *userRepository) FindUserInfoByUserName(userName string, password string) User {
	var recode User
	tx := ur.db.Select([]string{}).
		Where("user_name=?", userName)
	if password != "" {
		tx.Where("password=? AND status=?", password, 1)
	}
	tx.First(&recode)
	return recode
}

func (ur userRepository) CreateUser(userName string, password string, ip string) (User, error) {
	var user = User{
		UserName:  userName,
		Password:  password,
		SecretKey: util.MD5(strconv.FormatInt(util.Get(1), 10)),
		AppId:     util.Get(1),
		Email:     "1844066417@qq.com",
		Phone:     18689223002,
		Ip:        ip,
	}
	tx := ur.db.Create(&user)
	return user, tx.Error
}
