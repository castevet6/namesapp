package main

import (
    _ "fmt"
    "flag"
    "log"
    "net/http"
    "os"

    "github.com/castevet6/dev-academy-names-go/pkg/models/jsonreader"
    "github.com/joho/godotenv"
)

// Type application is container struct collecting the dependencies under a single wrapper.
//   devMode = boolean flag to indicate we are in development mode (CORS enabled)
//   infoLog = logger instance
//   names = data model wrapping the connection to data source file/database
type application struct {
    devMode     bool
    infoLog     *log.Logger
    names       *jsonreader.NameModel
}

func main() {
    // load environment variables
    godotenv.Load()

    // extract port address and datasource path from environment variables
    addr := ":" + os.Getenv("PORT")
    datasource := os.Getenv("NAMES_APP_DATA_PATH")

    // resolve cmd line flags
    devMode := flag.Bool("devmode", false, "Development mode configured (allow CORS)")
    flag.Parse()

    // set up errorlog
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // set up dependencies in application instance
    app := &application{
        devMode:    *devMode,
        infoLog:    infoLog,
        names:      jsonreader.New(datasource),  
    }

    router := app.routes()
    spa := spaHandler{ staticPath: "ui", indexPath: "index.html" }
    router.PathPrefix("/").Handler(spa)

    // start the server
    app.infoLog.Printf("Dev mode enabled: %v\n", app.devMode)
    if app.devMode == false {
        app.infoLog.Printf("  start with flag -devmode to enable\n")
    }
    app.infoLog.Printf("Start server on address %s\n", addr)
    err := http.ListenAndServe(addr, router)
    errorLog.Fatal(err)
} 
