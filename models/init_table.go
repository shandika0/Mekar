package models

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {

	// db.Debug().DropTableIfExists(&User{}, &Pekerjaan{}, &Pendidikan{})
	db.Debug().AutoMigrate(&User{}, &Pekerjaan{}, &Pendidikan{})
}
