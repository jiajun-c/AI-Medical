package crud

import (
	"ai-medical/model"
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
