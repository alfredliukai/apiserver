package model

import (
	"apiserver/pkg/auth"
	"gopkg.in/go-playground/validator.v9"
	"apiserver/pkg/constvar"
	"fmt"
	"sync"
)

type UserModel struct {
	BaseModel
	Username	string	`json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password	string	`json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

type UserInfo struct {
	Id			uint64	`json:"id"`
	Username	string	`json:"username"`
	SayHello	string	`json:"sayHello"`
	Password	string	`json:"password"`
	CreatedAt	string	`json:"createdAt"`
	UpdatedAt	string	`json:"updatedAt"`
}

type UserList struct {
	Lock *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}

func(u *UserModel) TableName() string{
	return "tb_users"
}


func(u *UserModel) Create() error{
	return DB.Self.Create(&u).Error
}

func Delete(id uint64)error{
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func(u *UserModel)Update() error{
	return DB.Self.Save(u).Error
}

func GetUser(username string)(*UserModel,error){
	u :=&UserModel{}
	d :=DB.Self.Where("username=?",username).First(&u)
	return u,d.Error
}

func ListUser(username string, offset uint, limit uint)([]*UserModel,uint64,error){
	if limit == 0{
		limit = constvar.DefaultLimit
	}
	var count uint64
	users := make([]*UserModel,0)

	where := fmt.Sprintf("username like '%%%s%%'",username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error;err !=nil{
		return users,count,err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error;err !=nil{
		return users,count,err
	}
	return users,count,nil
}

func(u *UserModel) Compare(pwd string)(err error){
	err = auth.Compare(u.Password,pwd)
	return
}

func(u *UserModel) Encrypt()(err error){
	u.Password,err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}