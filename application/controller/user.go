package controller

import (
	"github.com/shun198/go-crm/config"
	"github.com/shun198/go-crm/model"
	"github.com/shun198/go-crm/serializer"
	"gorm.io/gorm"
)

type UserController struct {
	Env *config.Env
}

var DB *gorm.DB

func (uc UserController) GetUsers() *serializer.ListResponse {
	var users []model.User
	return &serializer.ListResponse{
		Results: users,
	}
}
