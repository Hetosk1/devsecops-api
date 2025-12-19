package main 

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWrite, r *http.Request){
		w.Write([]byte("API is running"))
	})

	log.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
