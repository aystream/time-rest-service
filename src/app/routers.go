package app

import (
	"github.com/aystream/time-rest-service/src/app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Логирование внутренних запросов
func RESTLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Маршрут - подтип который содержит информацию по запросу
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Queries     []string
}

type Routes []Route

var routes = Routes{}

// использует обменные интерфейсы и возвращает новый routers
func (a *App) NewRouter() {
	router := mux.NewRouter().StrictSlash(true)

	routes = Routes{
		Route{
			Name:        "GetTimeNow",
			Method:      "GET",
			Path:        "/time/now",
			HandlerFunc: a.GetTimeNow,
		},
		Route{
			Name:        "GetTimeNowString",
			Method:      "GET",
			Path:        "/time/string",
			HandlerFunc: handlers.GetTimeNowString,
			Queries:     []string{"time", "{time}"},
		},
		Route{
			Name:        "AddTime",
			Method:      "GET",
			Path:        "/time/add",
			HandlerFunc: handlers.AddTime,
			Queries:     []string{"time", "{time}", "delta", "{delta}"},
		},
		Route{
			Name:        "SetTime",
			Method:      "POST",
			Path:        "/time/set",
			HandlerFunc: a.SetTime,
		},
		Route{
			Name:        "ResetTime",
			Method:      "POST",
			Path:        "/time/reset",
			HandlerFunc: a.ResetTime,
		},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = RESTLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler).
			Queries(route.Queries...)
	}
	a.Router = router
}

func (a *App) GetTimeNow(w http.ResponseWriter, r *http.Request) {
	handlers.GetTimeNow(a.DB, w, r)
}
func (a *App) SetTime(w http.ResponseWriter, r *http.Request) {
	handlers.SetTime(a.DB, w, r)
}
func (a *App) ResetTime(w http.ResponseWriter, r *http.Request) {
	handlers.ResetTime(a.DB, w, r)
}
