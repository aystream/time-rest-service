package models

import (
	"strconv"
	"time"
)
import _ "time"

const (
	formatLayoutFloat64 = "060102.150405.000"
	formatLayoutString  = "06.01.02 15:04:05 00000"
)

type ServerTimeFloat64 struct {
	Time float64 `json:"time"`
}

// Преобразование формат в формат времени работы сервера
func (t *ServerTimeFloat64) FormatToTime() (*ServerTime, error) {
	timeFloat64String := strconv.FormatFloat(t.Time, 'f', 9, 64)

	millisecond := getMillisecondsByStringTime(timeFloat64String)
	stringTime := string(([]rune(timeFloat64String))[:len(timeFloat64String)-3])
	serverTime, e := time.Parse(formatLayoutFloat64, stringTime+"."+millisecond)

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
	millisecond := getMillisecondsByStringTime(timeFloat64String)
	stringTime := string(([]rune(timeFloat64String))[:len(timeFloat64String)-4])
	timeFloat64, _ := strconv.ParseFloat(stringTime+millisecond, 64)
	return &ServerTimeFloat64{timeFloat64}
}

// Получение милисекунд из даты строки (так как шаблон в go .000 для них)
func getMillisecondsByStringTime(stringTime string) string {
	return string(([]rune(stringTime))[len(stringTime)-3 : len(stringTime)])
}
