package app

import (
	"log"
	"main/internal/delivery"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", delivery.Home)
	mux.HandleFunc("/create", delivery.CreateAscii)
	mux.HandleFunc("/download", delivery.DownloadFile)
	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
