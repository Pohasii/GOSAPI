package main

// Импортируем необходимые зависимости. Мы будем использовать
// пакет из стандартной библиотеки и пакет от gorilla
import (
	"net/http"
	"github.com/gorilla/mux"
	"GOSAPI/api/status"
	"GOSAPI/api/product"
	"GOSAPI/src/api/AddFeedbackHandler"
	"github.com/gorilla/handlers"
	"os"
	"GOSAPI/api/faq"
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

	r.Handle("/status", status.StatusHandler).Methods("GET")
	r.Handle("/products", product.ProductsHandler).Methods("GET")
	r.Handle("/products/{slug}/feedback", test).Methods("POST")
	r.Handle("/faq", faq.Faq).Methods("GET")

	// Наше приложение запускается на 3000 порту.
	// Для запуска мы указываем порт и наш роутер
	// Обернем наш роутер функцией LoggingHandler.
	// Таким образом нам будет доступен каждый запрос.
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}