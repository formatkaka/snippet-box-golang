package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config   *Config
}

func getConfig() *Config {
	config := new(Config)

	flag.StringVar(&config.Addr, "addr", ":4000", "Server Port Address")
	flag.StringVar(&config.StaticDir, "static-dir", "./ui/static", "Path to static assets")

	return config
}

func main() {

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logError := log.New(os.Stdout, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)
	config := getConfig()
	app := &application{
		errorLog: logError,
		infoLog:  logInfo,
		config:   config,
	}

	server := &http.Server{
		Addr:     app.config.Addr,
		ErrorLog: logError,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting on port %s", config.Addr)
	err := server.ListenAndServe()
	app.errorLog.Fatal(err)

}
