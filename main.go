package main

import (
	"bbs-forgo/database"
	"bbs-forgo/entity/po"
	"bbs-forgo/log"
	"bbs-forgo/middleware"
	"bbs-forgo/routers"
	"github.com/gin-gonic/gin"
)

func DataBaseInit() error {
	//数据库初始化
	err := database.Conn()
	if err != nil {
		return err
	}
	err = database.GetConn().AutoMigrate(
		&po.User{},
		&po.UserInfo{},
	)
	if err != nil {
		return err
	}
	return nil
}

func RouterInit(e *gin.Engine) {
	routers.Login(e)
}

func MiddlewareInit(e *gin.Engine) {
	e.Use(middleware.GinLogger())
	e.Use(middleware.GinRecovery(true))
}

func main() {
	gin.ForceConsoleColor()
	log.InitLogger("debug")
	defer log.GetSugarLogger().Sync()
	defer log.GetLogger().Sync()
	r := gin.Default()
	//加载中间件
	MiddlewareInit(r)

	//加载路由
	RouterInit(r)

	//数据库初始化
	err := DataBaseInit()
	if err != nil {
		log.GetSugarLogger().Error("数据库初始化失败", err.Error())
		return
	}

	//CasbinInit()
	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
