package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const port = "8080"
	servMux := http.NewServeMux()
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}
	servMux.Handle("/app/", http.StripPrefix("/app/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	servMux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	servMux.HandleFunc("GET /api/healthz", handlerReadiness)
	servMux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	servMux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	servMux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	server := http.Server{
		Addr:    ":" + port,
		Handler: servMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
