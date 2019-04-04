package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func openDB(username,password,addr,name string) *gorm.DB{
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	fmt.Println(config)
	db,err:=gorm.Open("mysql",config)
	if err != nil{
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	setUpDB(db)
	return db
}

func setUpDB(db *gorm.DB){
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(1)
}

func InitSelfDB() *gorm.DB{
	return openDB(viper.GetString("db.username"),
		          viper.GetString("db.password"),
			      viper.GetString("db.addr"),
		          viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB{
	return InitSelfDB()
}

func(db *Database) Init(){
	DB = &Database{
		Self :GetSelfDB(),
	}
}

func(db *Database) Close(){
	DB.Self.Close()
}