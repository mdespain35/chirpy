package main

import (
	"log"
	"net/http"
)

func main() {
	servMux := http.NewServeMux()
	servMux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	servMux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	servMux.HandleFunc("/healthz", healthHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: servMux,
	}

	log.Fatal(server.ListenAndServe())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte{'O', 'K'})

}
