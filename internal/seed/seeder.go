package seed

import (
	"log"

	"github.com/Raybird/whale/internal/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

// Load ...
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
