package server

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gangjun06/book-server/middlewares"
	dbmodels "github.com/gangjun06/book-server/models/db"
	v1 "github.com/gangjun06/book-server/routes/v1"
	"github.com/gangjun06/book-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	etcInit()
	applyConfig()
	initDB()
	startServer()
}

func etcInit() {
	rand.Seed(time.Now().Unix())
}

func applyConfig() {
	config := utils.LoadConfig()
	utils.SetConfig(config)
}

func startServer() {
	config := utils.GetConfig()
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.LoadHTMLGlob("public/verify.html")
	r.Use(middlewares.Cors())

	version1 := r.Group("/v1")
	v1.InitRoutes(version1)

	r.Run(":" + config.Port)

}

func initDB() {
	config := utils.GetConfig()
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.SetDB(db)
	log.Print("Successed To Connect Database")
	log.Print("Connection To " + connectionInfo)

	log.Print("Performing AutoMigrate...")
	var models = []interface{}{&dbmodels.User{}}
	db.AutoMigrate(models...)
	log.Print("Successfully performed AutoMigrate")
}
