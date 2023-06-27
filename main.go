package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/handler"
	"github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/repository"
	"github.com/farhaanfsk/Go-Gorilla-Mux-CRUD/service"
	"github.com/gorilla/mux"
)

var (
	dbName = flag.String("dbname", "Employee", "Database name")
	dbHost = flag.String("dbhost", "127.0.0.1", "Database host")
	dbPort = flag.String("dbport", "3000", "Port for the database")
	dbUser = flag.String("dbuser", "postgres", "Database user")
	dbPass = flag.String("dbpass", "postgres", "Database password")
)

func main() {
	rep := repository.NewRepo(InitDb(*dbHost, *dbPort, *dbUser, *dbPass, *dbName))
	s := service.NewService(rep)
	h := handler.NewHandler(&s)
	router := mux.NewRouter()
	router.HandleFunc("/employees", h.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", h.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", h.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", h.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees", h.UpdateEmployee).Methods("PUT")
	http.ListenAndServe(":1234", router)
}

func InitDb(dbHost, dbPort, dbUser, dbPass, dbName string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database connection error", err.Error())
	}
	return db
}
