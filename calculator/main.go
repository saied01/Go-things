package main

import (
	"Goculate/api"
	"log/slog"
	"net/http"
	"os"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := http.NewServeMux()

	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/add", api.HandleAdd)
	router.HandleFunc("/sub", api.HandleSub)
	router.HandleFunc("/mult", api.HandleMult)
	router.HandleFunc("/div", api.HandleDiv)

	srv := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	logger.Info("running on :8081")
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("server failed", "err", err)
	}
}
