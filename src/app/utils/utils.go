package utils

import (
	"errors"
	"github.com/aystream/time-rest-service/src/app/models"
	"math"
	"time"
)

// Добавление продолжительности в наше время
func AddDurationInTimeByFloat64(duration float64, currentTime *models.ServerTime) (*models.ServerTime, error) {
	if duration >= 1000000.0 {
		return nil, errors.New(InvalidDurationFormat)
	}

	multiplier := 10000.0

	var parts [7]int
	for i := 0; i < 7; i++ {
		parts[i] = int(duration / multiplier)
		if i == 5 {
			duration = math.Round((duration-float64(parts[i])*multiplier)*1000) / 1000
			// для учета милисекунд дальше нам нужно 3 знака
			multiplier /= 1000
		} else {
			duration = math.Round((duration-float64(parts[i])*multiplier)*100) / 100
			multiplier /= 100
		}
	}

	year, month, day := currentTime.Time.Date()
	hour, min, sec := currentTime.Time.Clock()
	nanosecond := currentTime.Time.Nanosecond()

	newTime := time.Date(year+parts[0], month+time.Month(parts[1]), day+parts[2], hour+parts[3], min+parts[4],
		sec+parts[5], int(nanosecond)+parts[6]*1000, currentTime.Time.Location())

	currentTime.Time = newTime
	return currentTime, nil
}
