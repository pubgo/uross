package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pubgo/g/xconfig/xconfig_web"
	"github.com/pubgo/g/xdi"
	"github.com/pubgo/uross/uross/controller"
	"github.com/pubgo/uross/uross/models"
)

type App struct {
	xdi.In
	web         *xconfig_web.XWeb
	exampleCtrl controller.IExampleCtrl
	exampleMdl  models.IExampleModel
}

func (t App) InitRoutes() {
	web := t.web

	example := web.Group("/example")
	examples := web.Group("/examples")

	exampleCtrl := t.exampleCtrl
	example.
		POST("/", exampleCtrl.CreateOne).
		DELETE("", exampleCtrl.Delete)

	examples.
		POST("/", exampleCtrl.CreateMany)

	// https://github.com/d5/tengo
	// https://github.com/yuin/gopher-lua
	web.POST("/rpc/:name", func(ctx *gin.Context) {
		// example_id=123
		// comment_user=11111111122
		// 聚合上面两个接口查出来的数据，然后统一的回复
		t.exampleMdl.FindOne()
		t.exampleMdl.FindOne()
	})
}
