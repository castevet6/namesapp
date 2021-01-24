package main

import (
    "net/http"
    "os"
    "path/filepath"

    "github.com/gorilla/mux"
)

// static file server exmaple implemenation from gorilla/mux documentation
type spaHandler struct {
    staticPath  string
    indexPath   string
}

func (h spaHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    path, err := filepath.Abs(r.URL.Path)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    path = filepath.Join(h.staticPath, path)

    _, err = os.Stat(path)
    if os.IsNotExist(err) {
        // file does not exist, serve index.html
        http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// helper function which sets the routing with gorilla/mux
func (app *application) routes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/api/v1/names", app.returnAllNames).Methods("GET")
    router.HandleFunc("/api/v1/names/{name}", app.returnByName).Methods("GET")

    return router
}

