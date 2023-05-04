package controllers

import (
	"backend/models"
	"testing"
)

func TestGetUser(t *testing.T) {
	user, err := models.GetUser(1)
	if err != nil {
		t.Errorf("GetUser() = %s; want nil", err)
	}
	if user.Username != "hughdazz" {
		t.Errorf("GetUser() = %s; want hughdazz", user.Username)
	}
}
