package router

import (
	"ai-medical/api/handler"
	"ai-medical/middlewares"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	auth := middlewares.Auth()
	engine.POST("/signup", handler.Register)
	engine.POST("/signin", auth.LoginHandler)
	engine.GET("/AllRecords", handler.GetAllRecords)
	engine.GET("/OneRecord", handler.GetOneRecord)
	engine.POST("/DeleteOneRecord", handler.DeleteOneRecord)
	engine.POST("/analyze", handler.Analyze)
}
