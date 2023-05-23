package handler

import (
	"ai-medical/client"
	"ai-medical/crud"
	"ai-medical/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
	//"github.com/spf13/viper"
)

func Analyze(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	//_ := viper.GetString("storage.url") + file.Filename
	//content, err := file.Open()
	//
	//client.GetResult(file.Filename,))
	username := ctx.PostForm("username")
	src, err := file.Open()
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "upload file is broken",
		})
		return
	}
	//ctx.SaveUploadedFile()
	content := make([]byte, file.Size)
	_, err = src.Read(content)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "upload file is broken",
		})
		return
	}
	//PointRes, ArrowRes, err := client.GetResult(file.Filename, content)
	feat, err := client.GetResult(file.Filename, content)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "upload file is broken",
		})
		return
	}
	PointRes, ArrowRes := feat.Pointimage, feat.Arrowimage
	basePath := viper.GetString("storage.path")
	filename := strings.TrimSuffix(file.Filename, ".jpg")
	filename0 := filename + ".jpg"
	filename1 := filename + "_point.jpg"
	filename2 := filename + "_arrow.jpg"
	path := basePath + file.Filename
	path1 := basePath + filename1
	path2 := basePath + filename2
	os.WriteFile(path1, PointRes, 0644)
	os.WriteFile(path2, ArrowRes, 0644)
	os.WriteFile(path, content, 0644)

	baseUrl := viper.GetString("storage.url")
	url1 := baseUrl + filename1
	url2 := baseUrl + filename2
	url0 := baseUrl + filename0
	var poss []*crud.Pos
	for i := 0; i < 8; i++ {
		poss = append(poss, &crud.Pos{
			X: feat.Position[i*2],
			Y: feat.Position[i*2+1],
		})
	}
	points, _ := json.Marshal(poss)

	crud.InsertRecord(&model.Record{
		Username:  username,
		Time:      time.Now().Format(time.UnixDate),
		OriginURL: url0,
		PointURL:  url1,
		ArrowURL:  url2,
		PT:        feat.PT,
		MT:        feat.MT,
		TL:        feat.TL,
		Points:    points,
	})

	ctx.JSON(200, gin.H{
		"pointImg":  url1,
		"arrowImg":  url2,
		"originUrl": url0,
		"PT":        feat.PT,
		"MT":        feat.MT,
		"TL":        feat.TL,
		"pos":       poss,
	})

}
