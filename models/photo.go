package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" form:"title" json:"title" `
	Caption  string `gorm:"not null" form:"caption" json:"caption" valid:"required~Caption is required"`
	PhotoUrl string `gorm:"not null" form:"photo_url" json:"photo_url" valid:"required~Photo Url is required"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCode := govalidator.ValidateStruct(p)

	if errCode != nil {
		err = errCode
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
