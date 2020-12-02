package models

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {

	// db.Debug().DropTableIfExists(&User{}, &Pekerjaan{}, &Pendidikan{})
	db.Debug().AutoMigrate(&User{}, &Pekerjaan{}, &Pendidikan{})
	// db.Model(&User{}).AddForeignKey("pekerjaan_id", "tb_pekerjaan(id)", "CASCADE", "CASCADE")
	// db.Model(&User{}).AddForeignKey("pendidikan_id", "tb_pendidikan(id)", "CASCADE", "CASCADE")

}
