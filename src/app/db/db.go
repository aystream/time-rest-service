package db

import (
	"sync"
	"time"
)

// Псевдо база данных для одной записи текущего состояния сервера
type DB struct {
	mutex    sync.Mutex    // симофор синхронизации потоков
	Duration time.Duration // разница от текущего времени
}

// Сохранение разницу во времени
func (db *DB) SaveDuration(duration time.Duration) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.Duration = duration
}
