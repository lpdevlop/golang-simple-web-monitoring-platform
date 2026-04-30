package templetes

import (
	"fmt"
	"html/template"
	"net/http"
	"webscaper/db"
	"webscaper/hashing"
	"webscaper/models"
	"webscaper/scrapping"
)

var tmpl = template.Must(template.ParseFiles("templetes/index.html"))

func homeHandleer(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		url := r.FormValue("url")

		fmt.Println("URL from form:", url)

		pagetitle, body := scrapping.StartScraping(url)

		titlehash, bodyHash := hashing.CheckHash(body, pagetitle)

		record := models.Monitor{
			HashValuesTitle: titlehash,
			HashValuesBody:  bodyHash,
			Title:           pagetitle,
			Url:             url,
		}

		db.DB.Create(&record)

		views, _ := db.BuildComparison(url, titlehash, bodyHash)

		data := struct {
			Records []models.MonitorView
		}{
			Records: views,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func InitHTML() {
	http.HandleFunc("/", homeHandleer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
