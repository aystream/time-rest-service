package app

import (
	"github.com/aystream/time-rest-service/src/app/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Приложение
type App struct {
	Router *mux.Router //экземпляры маршрутизатора
	DB     *db.DB
}

// Инициализация приложение
func (a *App) Initialize() {
	a.DB = &db.DB{}
	a.NewRouter()
}

// Запустите приложение по определенном порту
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
