package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gangjun06/book-server/models"
)

func LoadConfig() *models.Config {
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var config models.Config
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Successfully Opened config.json")
	return &config
}
