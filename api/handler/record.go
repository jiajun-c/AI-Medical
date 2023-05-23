package handler

import (
	"ai-medical/crud"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllRecords(ctx *gin.Context) {
	username := ctx.Query("username")
	records, err := crud.GetAllRecords(username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, records)
}

func GetOneRecord(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	record, err := crud.GetOneRecord(id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, record)
}

func DeleteOneRecord(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	err := crud.DeleteRecord(id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
