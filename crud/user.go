package crud

import (
	"ai-medical/model"
	"errors"
	log "github.com/sirupsen/logrus"
)

func SearchUser(username string, password string) bool {
	get, err := Engine.Where("username = ? and password = ?", username, password).Get(&model.User{})
	if err != nil {
		log.Error("failed to search user", err)
		return false
	}
	return get
}

func RegisterUser(username string, password string) error {
	find := SearchUser(username, password)
	if find {
		return errors.New("user already registered")
	}
	_, err := Engine.InsertOne(model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}
