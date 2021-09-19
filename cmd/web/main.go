package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/formatkaka/snippet-box-golang/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr      string
	StaticDir string
	DSN       string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config   *Config
	snippets *mysql.SnippetModel
}

func getConfig() *Config {
	config := new(Config)

	flag.StringVar(&config.Addr, "addr", ":4000", "Server Port Address")
	flag.StringVar(&config.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&config.DSN, "dsn", "web:abcdefgh@/snippetbox?parseTime=true", "MySQL data source name")

	return config
}

func main() {

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logError := log.New(os.Stdout, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	config := getConfig()

	db, err := openDB(config)

	app := &application{
		errorLog: logError,
		infoLog:  logInfo,
		config:   config,
		snippets: &mysql.SnippetModel{DB: db},
	}

	if err != nil {
		app.errorLog.Fatal(err)
	}

	defer db.Close()

	server := &http.Server{
		Addr:     app.config.Addr,
		ErrorLog: logError,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting on port %s", config.Addr)
	err = server.ListenAndServe()
	app.errorLog.Fatal(err)

}

func openDB(config *Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", config.DSN)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
