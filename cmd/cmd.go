package cmd

import (
	"admin-demo/conf"
	"admin-demo/global"
	"admin-demo/router"
	"admin-demo/utils"
	"fmt"
)

func Start() {
	var initErr error
	//初始化配置
	conf.InitConf()
	//初始化日志
	global.Logger = conf.InitLogger()
	//初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	//初始化redis
	redisClient, initErr := conf.InitRedis()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.RedisClient = redisClient
	_ = global.RedisClient.Set("a", "a")
	_ = global.RedisClient.Set("b", "11")
	_ = global.RedisClient.Set("c", 11)
	a, _ := global.RedisClient.Get("a")
	b, _ := global.RedisClient.Get("b")
	c, _ := global.RedisClient.Get("c")
	fmt.Printf("%#v,%#v,%#v", a, b, c)

	//初始化异常统一抛出
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}

	//初始化路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("===clean===")
}
