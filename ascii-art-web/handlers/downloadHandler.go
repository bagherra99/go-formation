package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom du fichier à télécharger depuis les paramètres de requête
	filename := r.URL.Query().Get("filename")

	// Ouvrir le fichier à télécharger
	filePath := fmt.Sprintf("ascii-art-files/%s", filename)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Obtenir les informations sur le fichier pour récupérer sa taille
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Définir les en-têtes HTTP pour le téléchargement
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Envoyer le contenu du fichier en tant que réponse HTTP
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
