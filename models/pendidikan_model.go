package models

import "github.com/jinzhu/gorm"

type Pendidikan struct {
	gorm.Model
	Nama string `gorm:"not null"json:"nama"`
}

func (u *Pendidikan) TableName() string {
	return "tb_pendidikan"
}
