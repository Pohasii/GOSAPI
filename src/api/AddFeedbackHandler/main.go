package AddFeedbackHandler

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

// 'Это нужно переделать, сюда должен приходить айди продукта, отзыв и токен.
// тут не должно быть этой структуры и массива.
type product struct {
	Id int
	Name string
	Slug string
	Description string
}

var products = []product{
	product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters",
		Description : "Shoot your way to the top on 14 different hoverboards"},
	product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer",
		Description : "Explore the depths of the sea in this one of a kind"},
	product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park",
		Description : "Go back 65 million years in the past and ride a T-Rex"},
	product{Id: 4, Name: "Cars VR", Slug : "cars-vr",
		Description: "Get behind the wheel of the fastest cars in the world."},
	product{Id: 5, Name: "Robin Hood", Slug: "robin-hood",
		Description : "Pick up the bow and arrow and master the art of archery"},
	product{Id: 6, Name: "Real World VR", Slug: "real-world-vr",
		Description : "Explore the seven wonders of the world in VR"},
}

// 'Это нужно переделать, сюда должен приходить айди продукта, отзыв и токен.




// Этот хендлер позволяет добавить позитивный или негативный отзыв
// по конкретному продукту. Правильно было бы сохранить результат в базу
// данных, но для демо нам это не нужно.
var Run = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var product product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})
