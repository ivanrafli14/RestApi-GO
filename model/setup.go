package model

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host string
	Port string
	Database string
	Username string
	Password string
}

var DB *gorm.DB

func ConnectDatabase(){
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	configure := ConfigDB {
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	 dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	 	configure.Username, configure.Password, configure.Host, configure.Port, configure.Database)
	
	
	db, err := gorm.Open(mysql.Open(dsn) , &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	Migrate(db)
}

func Migrate(connection *gorm.DB){
	connection.Debug().AutoMigrate(
		&Song{},
		&Album{},
	)
}