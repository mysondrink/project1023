// 主程序入口
package main

import (
	"assemblyline/project1023/common"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// _ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

func main() {
	InitConfig()
	common.InitDB()
	// defer db.Close()

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	logFile := "gin.log"
	var f *os.File
	// Logging to a file.
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ = os.Create(logFile)
	} else {
		f, _ = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0666)
	}
	// gin.DefaultWriter = io.MultiWriter(f)

	defer f.Close()
	// 创建MultiWriter，包括文件和字符串缓冲区
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

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
