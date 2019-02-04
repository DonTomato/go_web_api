package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dontomato/go_web_api/db"
	"github.com/dontomato/go_web_api/global"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("Error with config file: %v\n", err)
		return
	}
	parser := json.NewDecoder(configFile)
	if err = parser.Decode(&global.Config); err != nil {
		fmt.Printf("Error with config file: %v\n", err)
		return
	}

	fmt.Printf("Config has been loaded: %v\n", global.Config)

	dbOk := db.CheckDbVersion()
	fmt.Printf("DB is ok: %v", dbOk)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", global.Config.AppPort), router))
}
