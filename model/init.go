package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Self *gorm.DB
	//Docker *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self:GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitSelfDB() *gorm.DB{
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}



func openDB(username,password,addr,name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db,err := gorm.Open("mysql",config)
	if err != nil {
		log.Errorf(err,"Database connection failed. Database name: %s", name)
	}
	setupDB(db)
	return db

}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)
}
