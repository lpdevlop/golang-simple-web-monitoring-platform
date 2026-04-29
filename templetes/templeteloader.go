package templetes

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templetes/index.html"))

func homeHandleer(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func InitHTML() {
	http.HandleFunc("/", homeHandleer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
