package router

import (
	_ "admin-demo/docs"
	"admin-demo/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type IFnRegisterRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRouters []IFnRegisterRouter
)

func RegisterRouter(fn IFnRegisterRouter) {
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

func InitRouter() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	//初始化gin框架，注册相关路由
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	//初始化基础路由
	initBasePlatformRouters()

	//注册自定义验证器
	registerCustomValidate()

	for _, fnRegisterRouter := range gfnRouters {
		fnRegisterRouter(rgPublic, rgAuth)
	}

	//集成swagger
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("server.port")
	if port == "" {
		port = "8889"
	}
	//err := r.Run(fmt.Sprintf(":%s", port))
	//if err != nil {
	//	panic(fmt.Sprintf("Start Server Err:%s", err.Error()))
	//}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}
	go func() {
		global.Logger.Infof(fmt.Sprintf("Start Server Listen:%s", port))
		//fmt.Println(fmt.Sprintf("Start Server Listen:%s", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf(fmt.Sprintf("Start Server Error:%s", err.Error()))
			//fmt.Println(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
	}()

	<-ctx.Done()
	//cancelCtx()
	ctx, cancelShutDown := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelShutDown()
	fmt.Println("===")
	fmt.Println(<-ctx.Done())
	fmt.Println("===111")

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf(fmt.Sprintf("Stop Server Error:%s", err.Error()))
		//fmt.Println(fmt.Sprintf("Stop Server Error:%s", err.Error()))
		return
	}

	global.Logger.Info("Stop Server Success")
	//fmt.Println("Stop Server Success")
}

func initBasePlatformRouters() {
	InitUserRouters()
}

func registerCustomValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}
			return false
		})
	}
}
