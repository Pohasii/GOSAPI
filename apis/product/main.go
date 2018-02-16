package product

import (
	"net/http"
	"encoding/json"
)

// ProductsHandler срабатывает в момент вызова роута /products
// Этот хендлер возвращает пользователю список продуктов для оценки.
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// Конвертируем наш слайс продуктов в json
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

// Создадим новый тип Product. Мы будем использовать
// этот тип для для хранения информации о опытах VR
type Product struct {
	Id int
	Name string
	Slug string
	Description string
}

// Создадим наш каталог VR экспериментов и сохраним его в виде слайса.
var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters",
		Description : "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer",
		Description : "Explore the depths of the sea in this one of a kind"},
	Product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park",
		Description : "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug : "cars-vr",
		Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood",
		Description : "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr",
		Description : "Explore the seven wonders of the world in VR"},
}