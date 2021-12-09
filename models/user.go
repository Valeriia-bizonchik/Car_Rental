package models

import (
	"errors"
	"github.com/Valeriia-bizonchik/CarRental/utils"
	"gorm.io/gorm"
)

type Role int

const (
	Customer Role = iota
	Admin
)

type User struct {
	gorm.Model `json:"-"`
	Role       Role   `json:"role"`
	Name       string `json:"name" gorm:"uniqueIndex"`
	Passwd     string `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Passwd == "" {
		err = errors.New("password can't be an empty")
		return
	}

	hPassword, err := utils.HashPassword(u.Passwd)
	if err != nil {
		return
	}

	u.Passwd = hPassword
	return
}
