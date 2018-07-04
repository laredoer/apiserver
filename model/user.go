package model

import "gopkg.in/go-playground/validator.v9"

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

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
