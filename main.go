package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleFunc)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущен на http://0.0.0.0:9000")
	http.ListenAndServe("0.0.0.0:9000", nil)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/html.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
