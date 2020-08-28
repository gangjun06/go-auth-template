package utils

import (
	"github.com/gangjun06/book-server/models"
	"github.com/jinzhu/gorm"
)

var (
	g_config *models.Config
	g_db     *gorm.DB
)

func SetConfig(config *models.Config) {
	g_config = config
}

func GetConfig() *models.Config {
	return g_config
}

func SetDB(db *gorm.DB) {
	g_db = db
}

func GetDB() *gorm.DB {
	return g_db
}
