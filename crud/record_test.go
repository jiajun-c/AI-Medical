package crud

import (
	"ai-medical/model"
	"encoding/json"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
	"xorm.io/xorm"
)

func initDB() {
	driver := "postgres"
	dsn := "postgres://admin:admin@127.0.0.1:25432/ai?sslmode=disable"
	log.Info("The db dsn:", dsn)
	Db, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		log.Error(err)
		panic("failed to initialize the database")
	}
	Engine = Db
	CheckTable(model.Record{})
}
func TestInsertRecord(t *testing.T) {
	initDB()
	var points []*Pos
	points = append(points, &Pos{
		X: 1.0,
		Y: 1.1,
	})
	js, _ := json.Marshal(points)
	err := InsertRecord(&model.Record{
		Username:  "cjj",
		Time:      time.Now().Format(time.UnixDate),
		OriginURL: "aa",
		PointURL:  "bb",
		ArrowURL:  "cc",
		Points:    js,
	})
	if err != nil {
		log.Error(err)
		t.Error(err)
	}
}

func TestGetAllRecords(t *testing.T) {
	initDB()
	records, err := GetAllRecords("cjj")
	if err != nil {
		t.Error(err)
		return
	}
	println(records[0].OriginURL)
}

func TestGetOneRecord(t *testing.T) {
	initDB()
	record, _ := GetOneRecord(1)
	println(string(record.Points))
	//println(string(data))
}
