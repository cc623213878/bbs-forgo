package main

import (
	"bbs-forgo/database"
	"bbs-forgo/log"
	"github.com/gin-gonic/gin"
)

func RouterInit(e *gin.Engine) {

}

func main() {
	gin.ForceConsoleColor()
	log.InitLogger("debug")
	defer log.GetSugarLogger().Sync()
	defer log.GetLogger().Sync()
	r := gin.Default()
	//加载中间件
	r.Use(log.GinLogger())
	r.Use(log.GinRecovery(true))
	err := database.Conn()
	if err != nil {
		log.GetSugarLogger().Error("数据库初始化失败", err.Error())
		return
	}
	//加载路由
	RouterInit(r)
	//CasbinInit()
	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
