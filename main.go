package main

import (
	"ClickPhonebook/logger"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"ClickPhonebook/handler"
	"html/template"
	"ClickPhonebook/db"
)

func main() {
	logger.NewLogger()
	// основные настройки к базе
	dsn := "dbuser:dbpassword@tcp(172.21.0.2:3306)/devdb?"
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
	r.HandleFunc("/contacts", handlers.List).Methods("GET")
	r.HandleFunc("/contacts/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/contacts/new", handlers.Add).Methods("POST")
	r.HandleFunc("/contacts/{id}/addphone", handlers.AddFormPhone).Methods("GET")
	r.HandleFunc("/contacts/{id}/addphone", handlers.AddPhone).Methods("POST")
	r.HandleFunc("/contacts/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/contacts/{id}", handlers.Update).Methods("POST")
	r.HandleFunc("/contacts/{id}", handlers.Delete).Methods("DELETE")


	//r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	//r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	//r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")



	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)

}