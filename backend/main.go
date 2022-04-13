package main

import (
	"d/go/routers"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// type spaHandler struct {
// 	staticPath string
// 	indexPath  string
// }

// // ServeHTTP inspects the URL path to locate a file within the static dir
// // on the SPA handler. If a file is found, it will be served. If not, the
// // file located at the index path on the SPA handler will be served. This
// // is suitable behavior for serving an SPA (single page application).
// func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// get the absolute path to prevent directory traversal
// 	path, err := filepath.Abs(r.URL.Path)
// 	if err != nil {
// 		// if we failed to get the absolute path respond with a 400 bad request
// 		// and stop
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// prepend the path with the path to the static directory
// 	path = filepath.Join(h.staticPath, path)

// 	// check whether a file exists at the given path
// 	_, err = os.Stat(path)
// 	if os.IsNotExist(err) {
// 		// file does not exist, serve index.html
// 		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
// 		return
// 	} else if err != nil {
// 		// if we got an error (that wasn't that the file doesn't exist) stating the
// 		// file, return a 500 internal server error and stop
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// otherwise, use http.FileServer to serve the static dir
// 	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
// }

func main() {
	//uncomment on first launch and comment after sucess
	// database.Create_basic_tables()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println("Path to exec: ", exPath)
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		// AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/").Subrouter()
	basic_auth := router.PathPrefix("/auth/").Subrouter()
	soc_auth := router.PathPrefix("/socialauth/").Subrouter()
	routers.Route_api(api)
	routers.Route_auth_basic(basic_auth)
	routers.Route_auth_social(soc_auth)
	// spa := spaHandler{staticPath: "../web", indexPath: "index.html"}
	// router.PathPrefix("/").Handler(spa)

	fmt.Println("Server is listening....")

	handler := cors.Handler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Port: ", port)
	server := http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	server.ListenAndServe()
}
