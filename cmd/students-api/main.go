package main

import (
	"github.com/nk-31012002/student-api/internal/config"
	"net/http"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//setup database
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welome to the student api"))
	})
	//setup server
}
