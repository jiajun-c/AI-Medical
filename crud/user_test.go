package crud

import (
	"ai-medical/model"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	initDB()
	CheckTable(model.User{})
	find := SearchUser("admin", "admin")
	if find {
		t.Error("register")
	}
}
