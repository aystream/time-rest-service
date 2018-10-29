package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/aystream/time-rest-service/src/app/models"
)

// Тестирование метода добавление ко времени временного промежутка
func TestAddDurationInTimeByFloat64(t *testing.T) {
	type args struct {
		duration    float64
		currentTime time.Time
	}

	// Наши тестовые запросы
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			args: args{
				duration:    020102.11203,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2020, time.Month(12), 25, 23, 24, 50, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    020102.0,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2020, time.Month(12), 25, 12, 04, 20, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    0,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    0.1,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2018, time.Month(11), 23, 22, 4, 20, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    1,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2018, time.Month(11), 24, 12, 4, 20, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    -020102.11203,
				currentTime: time.Date(2018, time.Month(11), 23, 12, 4, 20, 0, time.Now().Location()),
			},
			want: time.Date(2016, time.Month(10), 21, 00, 43, 50, 0, time.Now().Location()),
		},
		{
			args: args{
				duration:    99.99,
				currentTime: time.Date(2018, time.Month(01), 01, 0, 0, 0, 0, time.Now().Location()),
			},
			want: time.Date(2018, time.Month(04), 14, 03, 0, 0, 0, time.Now().Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddDurationInTimeByFloat64(tt.args.duration, &models.ServerTime{Time: tt.args.currentTime})
			if (err != nil) != tt.wantErr {
				t.Errorf("AddDurationInTimeByFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Time, tt.want) {
				t.Errorf("AddDurationInTimeByFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
