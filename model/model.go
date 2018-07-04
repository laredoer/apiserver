package model

import (
	"time"
	"sync"
)

type BaseModel struct {
	Id uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreateAt time.Time `gorm:"column:createAt" json:"-"`
	UpdateAt time.Time 	`gorm:"column:updateAt" json:"-"`
	DeleteAt *time.Time `gorm:"column:deleteAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
	SayHello string `json:"sayHello"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	UpdataAt string `json:"updateAt"`
}

type Userlist struct {
	lock *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
