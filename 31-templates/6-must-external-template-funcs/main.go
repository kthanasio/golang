package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(v string) string {
	return strings.ToUpper(v)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		t := template.Must(template.New("content.html").
			Funcs(template.FuncMap{"ToUpper": ToUpper}).
			ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{Nome: "Go", CargaHoraria: 40},
			{Nome: "Java", CargaHoraria: 80},
			{Nome: "Python", CargaHoraria: 120},
			{Nome: "JavaScript", CargaHoraria: 200},
			{Nome: "TypeScript", CargaHoraria: 180},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
