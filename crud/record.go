package crud

import (
	"ai-medical/model"
	log "github.com/sirupsen/logrus"
)

type Pos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func GetAllRecords(username string) ([]model.Record, error) {
	var records []model.Record
	err := Engine.Where("username = ?", username).Cols("id", "time").Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func InsertRecord(record *model.Record) error {
	_, err := Engine.InsertOne(record)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetOneRecord(id int) (*model.Record, error) {
	var record model.Record
	get, err := Engine.ID(id).Get(&record)
	if err != nil {
		return nil, err
	}
	if get == false {
		log.Error("fai")
	}
	return &record, nil
}

func DeleteRecord(id int) error {
	var record model.Record
	_, err := Engine.ID(id).Delete(&record)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
