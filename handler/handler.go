package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "_partials", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	product := []entity.Product{
		{ID: 1, Name: "MacBook Pro 2020", Price: 20000000, Stock: 1},
		{ID: 2, Name: "Asus VivoBook", Price: 8000000, Stock: 5},
		{ID: 3, Name: "MacBook Air 2020", Price: 16000000, Stock: 3},
	}
	data := map[string]interface{}{
		"title":   "Golang WEB",
		"content": product,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}

func ProductPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "_partials", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Product Page",
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini method Get"))
	case "POST":
		w.Write([]byte("Ini Method POST"))
	default:
		http.Error(w, "Error Is Happening", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "_partials", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, error := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "_partials", "layout.html"))

		if error != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		error = tmpl.Execute(w, data)

		if error != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening", http.StatusBadRequest)
}
