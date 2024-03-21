package main

import (
	"battle-brackets/db"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	r := mux.NewRouter()

	db := db.Init()

	port := os.Getenv("APP_PORT")
	http.ListenAndServe(":"+port, r)
}
