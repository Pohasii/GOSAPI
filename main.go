package main

// Импортируем необходимые зависимости. Мы будем использовать
// пакет из стандартной библиотеки и пакет от gorilla
import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"api/faq"
	"api/product"
	"api/AddFeedbackHandler"
	"api/status"
)

func main() {
	// Инициализируем gorilla/mux роутер
	r := mux.NewRouter()

	// Страница по умолчанию для нашего сайта это простой html.
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Статику (картинки, скрипти, стили) будем раздавать
	// по определенному роуту /static/{file}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	r.Handle("/status", Status.Run).Methods("GET")
	r.Handle("/products", Product.Run).Methods("GET")
	r.Handle("/products/{slug}/feedback", AddFeedbackHandler.Run).Methods("POST")
	r.Handle("/faq", Faq.Run).Methods("GET")

	// Наше приложение запускается на 3000 порту.
	// Для запуска мы указываем порт и наш роутер
	// Обернем наш роутер функцией LoggingHandler.
	// Таким образом нам будет доступен каждый запрос.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}