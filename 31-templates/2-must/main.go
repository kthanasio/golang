package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	cursoGo := Curso{
		Nome:         "GO",
		CargaHoraria: 180,
	}
	t := template.Must(template.New("CursoTemplate").
		Parse("Curso [{{.Nome}}] - Carga Horária [{{.CargaHoraria}}]"))
	err := t.Execute(os.Stdout, cursoGo)
	if err != nil {
		panic(err)
	}
}
