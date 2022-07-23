package main

import (
	orm "simpleRestAPI/dao"
	router "simpleRestAPI/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	defer orm.Close()
	//註冊資料庫
	if err := orm.InitSql(); err != nil {
		panic(err)
	}
	//cors
	r.Use(cors.Default())
	//註冊路由
	router.InitRouter(r)

	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
