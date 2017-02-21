package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, string("I would be index.html"))

}

func powerpost(res http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()
	fmt.Println("Post was called")
	fmt.Println(vars["SN"])
}

func main() {

	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://energymon@localhost:26257/powerread?sslmode=disable")
	if err != nil {
		log.Fatalf("error connection to the database: %s", err)
	}

	// Create the "powerread" table.
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS powerread (id INT PRIMARY KEY, watts INT)"); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/powerpost/", powerpost)

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
