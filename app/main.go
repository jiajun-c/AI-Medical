package main

import (
	"ai-medical/api/router"
	"ai-medical/crud"
	"ai-medical/middlewares"
	"ai-medical/model"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

func initConfig() {
	//path := "/home/bot/workspace/AI-Medical/config.yaml"
	path := "/config.yaml"
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("failed to load config file")
		panic(err)
	}
}

func initDB() {
	driver := viper.GetString("db.driver")
	dsn := viper.GetString("db.dsn")
	logrus.Info("The db dsn:", dsn)
	Db, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		logrus.Error(err)
		panic("failed to initialize the database")
	}
	crud.Engine = Db
	crud.CheckTable(model.User{})
}

func init() {
	initConfig()
	initDB()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())

	router.Route(r)
	addr := viper.GetString("api.addr")
	err := r.Run(addr)
	if err != nil {
		return
	}
}
