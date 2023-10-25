package main

import (
	"assemblyline/project1023/common"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

func main() {
	common.InitDB()
	// defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	// r.Run() // listen and serve on 0.0.0.0:8080
	panic(r.Run())
}
