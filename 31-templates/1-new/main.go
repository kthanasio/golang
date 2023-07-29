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

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso [{{.Nome}}] - Carga Horária [{{.CargaHoraria}}]")
	err := tmp.Execute(os.Stdout, cursoGo)
	if err != nil {
		panic(err)
	}
}
