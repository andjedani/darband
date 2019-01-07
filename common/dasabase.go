package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {

	// var dbPort string
	// var dbHost string
	// var postgresUser string
	// var postgresPassword string
	dbPort := viper.GetString("db_port")
	dbHost := viper.GetString("db_host")
	dbName := viper.GetString("db_name")
	postgresUser := viper.GetString("postgres_user")
	postgresPassword := viper.GetString("postgres_password")

	db, err := gorm.Open("postgres", "host="+dbHost+" port="+dbPort+" user="+postgresUser+" dbname="+dbName+" password="+postgresPassword)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
