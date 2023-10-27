// 主程序入口
package main

import (
	"assemblyline/project1023/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// _ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

func main() {
	InitConfig()
	common.InitDB()
	// defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	// r.Run() // listen and serve on 0.0.0.0:8080
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
