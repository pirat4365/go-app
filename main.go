package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-app/infrastructure"
	"go-app/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var database *gorm.DB

func init() {
	database = infrastructure.DatabaseConnect()
}

func main() {
	router := mux.NewRouter()

	migrate()
	router.HandleFunc("/", func(w http.ResponseWriter, router *http.Request) {
		fmt.Fprintf(w, "Добро пожаловать")
	})

	router.HandleFunc("/сreate-user", func(w http.ResponseWriter, router *http.Request) {
		user := models.Users{Name: "PETYNYA", Age: 25}

		database.Create(&user)
	})

	router.HandleFunc("/select-users", func(w http.ResponseWriter, router *http.Request) {
		var users []models.Users
		result := database.Find(&users)
		log.Print(result)

	})
	router.HandleFunc("/redirect", func(w http.ResponseWriter, router *http.Request) {
		http.Redirect(w, router, "", 302)
	})
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":80", router)
}

func migrate() {
	err := database.AutoMigrate(&models.Users{})
	if err != nil {
		log.Print(err)
	}
}
