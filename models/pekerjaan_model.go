package models

import "github.com/jinzhu/gorm"

type Pekerjaan struct {
	gorm.Model
	Nama string `gorm:"not null"json:"nama"`
}

func (u *Pekerjaan) TableName() string {
	return "tb_pekerjaan"
}
