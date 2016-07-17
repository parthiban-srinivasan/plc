package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "google.golang.org/appengine"
	"log"
	"net/http"
	_ "os"
	_ "time"
)

// var db *sql.DB
// var err error

// Init module...
func init() {

	        var db *sql.DB
	        var err error

	//      databaseConfig := os.Getenv("MYSQL_CONNECTION")


	//      db, err = sql.Open("mysql", databaseConfig)
	//      db, err = sql.Open("mysql", "root:root@tcp(104.196.22.179:3306)/testdb")

  //  Open validates the database arguments without creating connections
	db, err = sql.Open("mysql", "root@cloudsql(mygo-1217:us-central1:locdb)/testdb")

	if err != nil {
		log.Printf("not good")
		log.Fatal(err)
  }

  //  Test database connection
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

  //  Route request to appropriate handler
	http.HandleFunc("/", roothandler)

}
// end of Init function

// Root request will be handled.
func roothandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to County Service")
}
