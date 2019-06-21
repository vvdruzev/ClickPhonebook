package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"ClicksPhonebook/handler"
	"html/template"
	"ClicksPhonebook/db"
)

func main() {
	// основные настройки к базе
	dsn := "dbuser:dbpassword@tcp(172.17.0.2:3306)/devdb?"
	// указываем кодировку
	dsn += "&charset=utf8"
	// отказываемся от prapared statements
	// параметры подставляются сразу
	dsn += "&interpolateParams=true"
	mysqlrepo, err:=db.NewMysqlrepo(&dsn)

	db.SetRepository(mysqlrepo)

	mysqlrepo.Db.SetMaxOpenConns(10)

	err = mysqlrepo.Db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handlers := handler.NewHandler()
	handlers.Tmpl = template.Must(template.ParseGlob("templates/*"))

	// в целям упрощения примера пропущена авторизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	//r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	//r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	//r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")



	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)

}