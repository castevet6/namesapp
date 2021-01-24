package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

// helper function to set headers for enabling CORS
func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// returnAllNames gets all names via app.names handler, and returns corresponding JSON object.
func (app *application) returnAllNames(w http.ResponseWriter, r *http.Request) {
    namearray := app.names.GetAll()
    res, _ := json.Marshal(namearray)
    if app.devMode {
        app.infoLog.Println("CORS REQUEST")
        // DISABLED IN PRODUCTION MODE!!!
        enableCors(&w)
    }
    w.Write(res)
}

func (app *application) returnByName(w http.ResponseWriter, r *http.Request) {
    searchTerm := mux.Vars(r)["name"]
    name := app.names.GetName(searchTerm)
    res, _ := json.Marshal(name)

    if app.devMode {
        app.infoLog.Println("CORS REQUEST")
        // DISABLED IN PRODUCTION MODE!!!
        enableCors(&w)
    }
    w.Write(res)
    
}

