package common

import (
	"assemblyline/project1023/model"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// driverName := "mysql"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	// host := "localhost"
	// port := "3306"
	// database := "assembly_line"
	// username := "root"
	// password := "password"
	// charset := "utf8mb4"
	// loc := "Local"
	// user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, loc)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false, // skip the snake_casing of names
		},
	})
	// db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.UserTable{}, &model.Controller{}, &model.Car{}, &model.Toy{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
