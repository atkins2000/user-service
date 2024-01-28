package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInformation struct {
	Id      string `json:"id" gorm:"unique;default:gen_random_uuid()"`
	Name    string `json:"name"`
	Number  string `json:"number"`
	DOB     string `json:"birth_date"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

func (user *UserInformation) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.New().String()
	return
}
