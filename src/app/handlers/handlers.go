package handlers

import (
	"encoding/json"
	"github.com/aystream/time-rest-service/src/app/db"
	"github.com/aystream/time-rest-service/src/app/models"
	"github.com/aystream/time-rest-service/src/app/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ResponseError struct {
	Error string `json:"error"`
}

// RestJSONResponse выводит json-ответ
func RestJSONResponse(w http.ResponseWriter, r *http.Request, response interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}

// RestError выводит в консоль ошибку метода остановки и ошибку
func RestError(method string, err error) {
	log.Printf("REST метод %s: сервер не смог отправить ответ. ERROR: %s",
		method, err)
}

// получение времени сервера
func GetTimeNow(db *db.DB, w http.ResponseWriter, r *http.Request) {
	currentTimeServer := &models.ServerTime{Time: time.Now()}

	// учтем нашу дельту
	if db.Duration != 0 {
		currentTimeServer.Time = currentTimeServer.Time.Add(db.Duration)
	}
	convertToFormat := currentTimeServer.TimeToFloat64()

	err := RestJSONResponse(w, r, convertToFormat, http.StatusOK)
	if err != nil {
		RestError(r.Method, err)
	}
}

// получение представления времени в виде строки
func GetTimeNowString(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	getTime := vars["time"]

	serverTime, err := getServerTimeByTime(getTime)
	if err != nil {
		RestJSONResponse(w, r, &ResponseError{utils.InvalidDataFormat}, http.StatusBadRequest)
		return
	}

	err = RestJSONResponse(w, r, serverTime.ToString(), http.StatusOK)
	if err != nil {
		RestError(r.Method, err)
		return
	}
}

// Получить и проверить формат времени time
func getServerTimeByTime(time string) (*models.ServerTime, error) {
	timeFloat64, err := strconv.ParseFloat(time, 64)
	if err != nil {
		return nil, err
	}

	serverTimeFloat64 := &models.ServerTimeFloat64{Time: timeFloat64}
	serverTime, err := serverTimeFloat64.FormatToTime()

	if err != nil {
		return nil, err
	}
	return serverTime, nil
}

// увеличение или уменьшение значения времени
func AddTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	getTime := vars["time"]

	serverTime, err := getServerTimeByTime(getTime)
	if err != nil {
		RestJSONResponse(w, r, &ResponseError{utils.InvalidDataFormat}, http.StatusBadRequest)
		return
	}

	delta := vars["delta"]

	timeFloat64, err := strconv.ParseFloat(delta, 64)
	if err != nil {
		RestJSONResponse(w, r, &ResponseError{utils.InvalidDataFormat}, http.StatusBadRequest)
		return
	}
	// добавим Duration ко времени
	currentTime, err := utils.AddDurationInTimeByFloat64(timeFloat64, serverTime)
	if err != nil {
		RestJSONResponse(w, r, &ResponseError{utils.InvalidDataFormat}, http.StatusBadRequest)
		return
	}

	err = RestJSONResponse(w, r, currentTime.TimeToFloat64(), 200)
	if err != nil {
		RestError(r.Method, err)
	}
}

// корректировка времени сервера – без изменения времени сервера
func SetTime(db *db.DB, w http.ResponseWriter, r *http.Request) {
	getTime := r.FormValue("time")

	newTime, err := getServerTimeByTime(getTime)
	if err != nil {
		RestJSONResponse(w, r, &ResponseError{utils.InvalidDataFormat}, http.StatusBadRequest)
		return
	}

	// вычислим дельту
	delta := newTime.Time.Sub(time.Now())

	db.SaveDuration(delta)

	err = RestJSONResponse(w, r, "OK", http.StatusOK)
	if err != nil {
		RestError(r.Method, err)
		return
	}
}

// Сброс времени
func ResetTime(db *db.DB, w http.ResponseWriter, r *http.Request) {
	db.SaveDuration(time.Duration(0))
	err := RestJSONResponse(w, r, "OK", http.StatusOK)
	if err != nil {
		RestError(r.Method, err)
		return
	}
}
