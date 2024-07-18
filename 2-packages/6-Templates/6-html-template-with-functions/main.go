package main

import (
	"html/template"
	"os"
	"strings"
)

/*
	Introdução

	Em algum momento você já construiu uma aplicação dinâmica? Exemplo, precisar
	gerar um HTML dinâmico com os valores do banco de dados ou um e-mail com um
	código dinâmico

	O Golang possui um sistema de templates embutidos feitos para isso. Iremos fazer vários
	exemplos de templates, nesse arquivo iremos trabalhar com um arquivo HTML e listar
	dinamicamente os cursos com páginas dinâmicas e incluir uma função que deixará os títulos
	em caixa alta
*/

type Course struct {
	Name string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	//Em ParseFiles fazemos o destructuring dos templates
	tmp := template.New("content.html")
	tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
	tmp = template.Must(tmp.ParseFiles(templates...))

	err := tmp.Execute(os.Stdout, Courses{
		{"Golang forever", 40},
		{"Cursos da Web para Jr JavaScript Embarcar na área de TI em 6 meses ganhando 100 mil por minuto", 10},
		{"Curso de Reacto para ganhar muito dinheiro pq não tem ninguem que sabe esse framework", 40},
		{"PHP", 10},
		{"Javao Legado", 40},
		{"C# nao eh c# nao viu", 80},
	})
	if err != nil {
		panic(err)
	}
}