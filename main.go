package main

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"url_shorter/repository"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	repo := repository.UrlRepository{Cache: client}

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

	http.Redirect(w, r, url, http.StatusSeeOther)
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
