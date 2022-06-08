package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

type Livro struct {
	Id         int
	Nome       string
	Autor      string
	NumPaginas int
	Genero     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "estantelivros"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("ui/tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	// Abre a conexão com o banco de dados utilizando a função dbConn()
	db := dbConn()
	// Realiza a consulta com banco de dados e trata erros
	selDB, err := db.Query("SELECT * FROM livros ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Livro{}

	// Monta um array para guardar os valores da struct
	res := []Livro{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var id, numPaginas int
		var nome, autor, genero string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &nome, &autor, &numPaginas, &genero)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Nome = nome
		n.Autor = autor
		n.NumPaginas = numPaginas
		n.Genero = genero

		// Junta a Struct com Array
		res = append(res, n)
	}

	// Abre a página Index e exibe todos os registrados na tela
	tmpl.ExecuteTemplate(w, "Index", res)

	// Fecha a conexão
	defer db.Close()
}

func Show(w http.ResponseWriter, r http.Request) {
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT FROM livros WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	n := Livro{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
		var id, numPaginas int
		var nome, autor, genero string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &nome, &autor, &numPaginas, &genero)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Nome = nome
		n.Autor = autor
		n.NumPaginas = numPaginas
		n.Genero = genero
	}

	// Mostra o template
	tmpl.ExecuteTemplate(w, "Show", n)

	// Fecha a conexão
	defer db.Close()
}

func Edit(w http.ResponseWriter, r http.Request) {
	// Abre a conexão com banco de dados
	db := dbConn()

	// Pega o ID do parametro da URL
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT FROM livros WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Livro{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		//Armazena os valores em variaveis
		var id, numPaginas int
		var nome, autor, genero string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &nome, &autor, &numPaginas, &genero)
		if err != nil {
			panic(err.Error())
		}

		// Envia os resultados para a struct
		n.Id = id
		n.Nome = nome
		n.Autor = autor
		n.NumPaginas = numPaginas
		n.Genero = genero
	}

	// Mostra o template com formulário preenchido para edição
	tmpl.ExecuteTemplate(w, "Edit", n)

	// Fecha a conexão com o banco de dados
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {

	//Abre a conexão com banco de dados usando a função: dbConn()
	db := dbConn()

	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		nome := r.FormValue("nome")
		autor := r.FormValue("autor")
		numPaginas := r.FormValue("numPaginas")
		genero := r.FormValue("genero")

		// Prepara a SQL e verifica errors
		insForm, err := db.Prepare("INSERT INTO livros(nome, autor, num_paginas, genero) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Insere valores do formulario com a SQL tratada e verifica errors
		insForm.Exec(nome, autor, numPaginas, genero)

		// Exibe um log com os valores digitados no formulário
		log.Println("INSERT: Nome: " + nome + " | Autor: " + autor + " | NumPaginas: " + numPaginas + " | Genero: " + genero)
	}

	// Encerra a conexão do dbConn()
	defer db.Close()

	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}
