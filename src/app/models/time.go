package models

import (
	"strconv"
	"time"
)
import _ "time"

const (
	formatLayoutFloat64 = "060102.150405"
	formatLayoutString  = "06.01.02 15:04:05 00000"
)

type ServerTimeFloat64 struct {
	Time float64 `json:"time"`
}

// Преобразование формат в формат времени работы сервера
func (t *ServerTimeFloat64) FormatToTime() (*ServerTime, error) {
	stringTime := strconv.FormatFloat(float64(t.Time), 'f', 6, 64)
	serverTime, e := time.Parse(formatLayoutFloat64, stringTime)

	if e != nil {
		return nil, e
	}

	return &ServerTime{serverTime}, nil
}

type ServerTimeString struct {
	Time string `json:"str"`
}

type ServerTime struct {
	Time time.Time
}

// Преобразование времени в строку
func (t *ServerTime) ToString() *ServerTimeString {
	timeString := t.Time.Format(formatLayoutString)
	return &ServerTimeString{timeString}
}

// Преобразование времени в float64
func (t *ServerTime) TimeToFloat64() *ServerTimeFloat64 {
	timeFloat64String := t.Time.Format(formatLayoutFloat64)
	timeFloat64, _ := strconv.ParseFloat(timeFloat64String, 64)
	return &ServerTimeFloat64{timeFloat64}
}
