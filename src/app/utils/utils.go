package utils

import (
	"errors"
	"github.com/aystream/time-rest-service/src/app/models"
	"time"
)

// Преобразование float64 в duration
func AddDurationInTimeByFloat64(duration float64, currentTime *models.ServerTime) (*models.ServerTime, error) {
	if duration >= 1000000.0 {
		return nil, errors.New(InvalidDurationFormat)
	}
	currentDuration := time.Duration(0)

	currentDuration.Hours()
	multiplier := 10000.0

	var years, months, days, hours, minutes, seconds, milliseconds int
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			years = int(duration / multiplier)
			duration = duration - float64(years)*multiplier
		case 1:
			months = int(duration / multiplier)
			duration = duration - float64(months)*multiplier
		case 2:
			days = int(duration / multiplier)
			duration = duration - float64(days)*multiplier
		case 3:
			hours = int(duration / multiplier)
			duration = duration - float64(hours)*multiplier
		case 4:
			minutes = int(duration / multiplier)
			duration = duration - float64(minutes)*multiplier
		case 5:
			seconds = int(duration / multiplier)
			duration = duration - float64(seconds)*multiplier
		case 6:
			milliseconds = int(duration / multiplier)
			duration = duration - float64(milliseconds)*multiplier
		}

		multiplier = multiplier / 100
	}

	year, month, day := currentTime.Time.Date()
	hour, min, sec := currentTime.Time.Clock()

	newTime := time.Date(year+years, month+time.Month(months), day+days, hour, min, sec, int(currentTime.Time.Nanosecond()), currentTime.Time.Location())

	currentTime.Time = newTime
	return currentTime, nil
}
