package main

import (
	"encoding/json"
	"log"
	"net/http"
	"url_shorter/db"
	"url_shorter/repository"
)

func main() {
	database := db.CreateDb()
	defer database.Pool.Close()

	repo := repository.UrlRepository{Db: database}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /url", func(writer http.ResponseWriter, request *http.Request) {
		createUrl(writer, request, &repo)
	})
	mux.HandleFunc("GET /url/{id}", func(writer http.ResponseWriter, request *http.Request) {
		getUrl(writer, request, &repo)
	})

	addr := ":8080"

	log.Printf("Server start on %s\n", addr)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatal(err)
	}

}

func getUrl(w http.ResponseWriter, r *http.Request, repo repository.IUrlRepository) {
	url, err := repo.Find(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)
}

func createUrl(w http.ResponseWriter, r *http.Request, repo repository.IUrlRepository) {
	url, err := repo.Create(r.FormValue("url"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(url)
}
