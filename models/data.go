package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type File struct {
	GormModel
	FileName    string `gorm:"not null;uniqueIndex" form:"data"`
	Size        int    `gorm:"not null" form:"size"`
	ContentType string `gorm:"not null" form:"content-type"`
}

func (u *File) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
