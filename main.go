package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, router *http.Request) {
		fmt.Fprintf(w, "Добро пожаловать")
	})

	router.HandleFunc("/redirect", func(w http.ResponseWriter, router *http.Request) {
		http.Redirect(w, router, "", 302)
	})
	fmt.Println("Сервер запущен")
	http.ListenAndServe(":80", router)
}
