package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"FinalTaskAPI/helpers"
)

type User struct {
	GormModel
	Username string  `gorm:"uniqueIndex;not null" form:"username" json:"username" valid:"required~Username is required"`
	Email    string  `gorm:"uniqueIndex;not null" form:"email" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password string  `gorm:"not null" form:"password" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
