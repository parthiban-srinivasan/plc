package county

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
  "google.golang.org/appengine"
	_ "google.golang.org/appengine/log"
	"log"
	"net/http"
	_ "os"
	_ "time"
)

var db *sql.DB
var err error

// Init module...
func init() {

	//      databaseConfig := os.Getenv("MYSQL_CONNECTION")

	//      db, err = sql.Open("mysql", databaseConfig)
	//      db, err = sql.Open("mysql", "root:root@tcp(104.196.22.179:3306)/testdb")

	//  Open validates the database arguments without creating connections
	db, err = sql.Open("mysql", "root@cloudsql(mygo-1217:us-central1:locdb)/testdb")

	if err != nil {
		log.Printf("not good")
		log.Fatal(err)
	}

	//  Root request is handled here
	http.HandleFunc("/m", rootHandler)

	//  Health check is handled by "healthz" handler
	http.HandleFunc("/healthy", healthyHandler)

	//  Create table via createhandler
	http.HandleFunc("/create", createHandler)

	//  Warmup of instance (code load during instance creation) is handled by here
	http.HandleFunc("/_ah/warmup", warmupHandler)
}

// end of Init function

// Root request will be handled.
func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome to County Service")
}

// Warmup request will be handled here.
func healthyHandler(w http.ResponseWriter, r *http.Request) {

	//  Test database connection
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprint(w, "Healthy - county service")
	}
}

// Warmup request will be handled here.
func createHandler(w http.ResponseWriter, r *http.Request) {

	create_stmt := `CREATE TABLE IF NOT EXISTS zip (
				name       VARCHAR(15),
				state      CHAR(2)
			)`
	_, err := db.Exec(create_stmt)

	if err != nil {
		fmt.Fprint(w, "Failed - County table not created")
	} else {
		fmt.Fprint(w, "County table created")
	}
}

// Warmup request will be handled here.
func warmupHandler(w http.ResponseWriter, r *http.Request) {
  _ := appengine.NewContext(r)
	//log.Infof(ctx, "warmup done")
	fmt.Fprint(w, "Service warmed up")
}
