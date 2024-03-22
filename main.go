package main

import (
	"battle-brackets/controllers"
	"battle-brackets/db"
	"battle-brackets/repos"
	"battle-brackets/services"
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

	InitUserController(r, db)

	port := os.Getenv("APP_PORT")
	http.ListenAndServe(":"+port, r)
}

func InitUserController(router *mux.Router, db *gorm.DB) {
	r := repos.NewUserRepo(db)
	s := services.NewUserService(r)
	c := controllers.NewUserController(s)

	muxBase := "/api/users"
	router.HandleFunc(muxBase, c.Show).Methods(http.MethodGet)
	router.HandleFunc(muxBase, c.Create).Methods(http.MethodPost)
}
