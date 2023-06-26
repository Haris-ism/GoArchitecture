package postgre

import (
	"tes/models"
	"tes/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	postgreDB struct {
		postgre *gorm.DB
	}
	PostgreInterface interface {
		Insert(req models.ItemList) error
	}
)

func InitPostgre() PostgreInterface {
	host := utils.GetEnv("POSTGRE")
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Failed to Init Postgre, Err:", err)
	} else {
		logrus.Printf("Init Postgre Success")
	}
	db.AutoMigrate(&models.ItemList{})

	return &postgreDB{
		postgre: db,
	}
}
