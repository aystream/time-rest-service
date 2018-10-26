package db

import "time"

// Псевдо база данных для одной записи текущего состояния сервера
type DB struct {
	Duration time.Duration // разница от текущего времени
}

// Сохранение разницу во времени
func (db *DB) SaveDuration(duration time.Duration) {
	db.Duration = duration
}
