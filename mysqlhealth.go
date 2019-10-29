package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//App Application structure
// composed of the Router and Database string in the dsn format
type App struct {
	Router   *mux.Router
	Database string
}

//SetupRouter Sets up gorilla mux router endpoints
func (app *App) SetupRouter(healthEndpoint string) {
	app.Router.
		Methods("GET", "HEAD").
		Path(healthEndpoint).
		HandlerFunc(app.getHealth)
}

//getHealth handler function for the /health endpoint
func (app *App) getHealth(w http.ResponseWriter, r *http.Request) {
	connectionString := fmt.Sprintf("%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", app.Database)
	status := http.StatusServiceUnavailable
	text := []byte("")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		text = []byte("Failed to connect to mysql\n")
	} else {
		// we hsouldn't have more than 1 connection at a time for healthchecking
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(0)
		// Set max lifetime for a connection.
		db.SetConnMaxLifetime(1 * time.Minute)

		errPing := db.Ping()
		if errPing != nil {
			text = []byte("Failed to ping mysql\n")
		} else {
			_, err := db.Exec("SELECT 1;")
			if err != nil {
				text = []byte("Failed to query mysql\n")
			} else {
				status = http.StatusOK
				text = []byte("OK\n")
			}

		}
	}
	w.WriteHeader(status)
	w.Write(text)
	db.Close()
}

const (
	program           = "mysqlhealth"
	version           = "0.0.4"
	defaultDataSource = "mysql:mysql@tcp(localhost:3306)/"
)

var (
	versionF       = flag.Bool("version", false, "Print version information and exit.")
	listenAddressF = flag.String("web.listen-address", ":42005", "Address to listen on for web interface")   //TO be implemented
	telemetryPathF = flag.String("web.health-path", "/health", "Path under which to expose health endpoint") //TO be implemented
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s %s Exposes an http health endpoint for mysql health checks.\n", os.Args[0], version)
		fmt.Fprintf(os.Stderr, "It uses MYSQL_SOURCE_NAME for mysql connection environment variable with following format: https://github.com/go-sql-driver/mysql#dsn-data-source-name\n")
		fmt.Fprintf(os.Stderr, "Default value is %q.\n\n", defaultDataSource)
		fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionF {
		fmt.Println(version)
		os.Exit(0)
	}

	dsn := os.Getenv("MYSQL_SOURCE_NAME")
	if dsn == "" {
		dsn = defaultDataSource
	}

	app := &App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: dsn,
	}

	app.SetupRouter(*telemetryPathF)

	log.Fatal(http.ListenAndServe(*listenAddressF, app.Router))
}
