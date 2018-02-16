package Faq

import (
	"net/http"
	"encoding/json"
)

type maps struct {
	Name string
	Description string
}

// Создадим наш каталог VR экспериментов и сохраним его в виде слайса.
var list = []maps{
	maps{Name: "Hi",
		Description : "Hi Bogdan"},
}

var Run = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	payload, _ := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})