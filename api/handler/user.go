package handler

import (
	"ai-medical/crud"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context) {
	username, password := ctx.Query("username"), ctx.Query("password")
	err := crud.RegisterUser(username, password)
	if err != nil {
		ctx.JSON(401, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "successfully registered",
	})
}
