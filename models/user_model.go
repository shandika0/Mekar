package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nama         string     `gorm:"not null" json:"nama"`
	TanggalLahir string     `gorm:"not null" json:"tanggal_lahir"`
	NoKtp        int        `gorm:"not null" json:"no_ktp"`
	PekerjaanId  int        `gorm:"not null" json:"pekerjaan_id"`
	Pekerjaan    Pekerjaan  `gorm:"foreignkey:PendidikanID;references:id"`
	PendidikanId int        `gorm:"not null" json:"pendidikan_id"`
	Pendidikan   Pendidikan `gorm:"foreignkey:PendidikanID;references:id"`
}

func (u *User) TableName() string {
	return "tb_user"
}
