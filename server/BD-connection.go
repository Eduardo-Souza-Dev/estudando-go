package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Primeira inserção ao banco

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/biblioteca")
	if err != nil {
		panic(err.Error())
	}
	//O defer é usado para fechar uma função, nesse caso ele encerra a conexão com o banco de dados
	defer db.Close()
	insert, err := db.Query("INSERT INTO livros (titulo)  VALUES ('Livro de Teste')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Livro inserido com sucesso!")

	exportConnection(db)
}

func exportConnection(db *sql.DB) {
	fmt.Println(db)
}
