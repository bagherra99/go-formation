package main

import (
	handlers "AsciiArtWeb/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handlers.LoadTemplates()

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)

	fmt.Println("Serveur en cours d'exécution sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
