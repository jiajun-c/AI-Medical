package router

import (
	"ai-medical/middlewares"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	auth := middlewares.Auth()
	engine.POST("/register")
	engine.POST("/login", auth.LoginHandler)
}
