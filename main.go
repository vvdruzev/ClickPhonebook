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
	"log"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MysqlDB           string `envconfig:"MYSQL_DATABASE"`
	MysqlUser         string `envconfig:"MYSQL_USER"`
	MysqlPassword     string `envconfig:"MYSQL_PASSWORD"`
	ProxyUrl			 string `envconfig:"HTTP_PROXY"`
}

func main() {
	logger.NewLogger()
	// основные настройки к базе
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Error(err)
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?",cfg.MysqlUser,cfg.MysqlPassword,cfg.MysqlDB)
	// указываем кодировку
	dsn += "&charset=utf8"
	// отказываемся от prapared statements
	// параметры подставляются сразу
	dsn += "&interpolateParams=true"
	mysqlrepo, err:=db.NewMysqlrepo(&dsn)
	if err !=nil  {
		logger.Error("Error DB. Please check your connect for DB",err,dsn)
		log.Fatal()
	}

	db.SetRepository(mysqlrepo)

	mysqlrepo.Db.SetMaxOpenConns(10)

	err = mysqlrepo.Db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	logger.Info("Connect to DB ",cfg.MysqlDB, ", user " , cfg.MysqlUser)

	handlers := handler.NewHandler()
	handlers.Tmpl = template.Must(template.ParseGlob("/templates/*"))

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
	r.HandleFunc("/search", handlers.Search).Methods("GET")


	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)

}