package model

import (
	"gopkg.in/go-playground/validator.v9"
	"apiserver/pkg/constvar"
	"fmt"
	"apiserver/pkg/auth"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string{
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func DeleteUser(id uint64) error{
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}
func ListUser(username string, offset, limit int) ([]*UserModel,uint64,error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel,0)

	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'",username)

	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users,count,err
	}
	return users,count,nil

}

func (u *UserModel) Compare(pwd string) (err error){
	err = auth.Compare(u.Password,pwd)
	return
}

func (u *UserModel) Encrypt() (err error){
	u.Password,err =auth.Encrypt(u.Password)
	return
}



func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
