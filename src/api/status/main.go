package Status

import (
	"net/http"
	// "fmt"
	"encoding/json"
)

// Необходимо реализовать хендлер NotImplemented.
// Этот хендлер просто возвращает сообщение "Not Implemented"

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// Конвертируем наш слайс продуктов в json
	payload, _ := json.Marshal("Not Implemented")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}) // удалить - уже не нужно

// Хендлер StatusHandler будет срабатывать в тот момент момент, когда
// пользователь обращается по роуту /status. Этот хендлер просто возвращает
// строку "API is up and running".
var Run = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	payload, _ := json.Marshal("API is up and running")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})