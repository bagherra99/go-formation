package handlers

import (
	"AsciiArtWeb/src"
	"AsciiArtWeb/src/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type FormData struct {
	Text     string
	Banner   string
	Results  []string
	Filename string
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	data400 := struct {
		Title   string
		Message string
	}{
		Title:   "Erreur 400 - Requête incorrecte",
		Message: "Mauvaise requête ou Caractères non-ASCII détectés",
	}

	// Ajouter le nom du fichier dans la structure de données du formulaire
	// formData := FormData{
	// 	Text:     text,
	// 	Banner:   banner,
	// 	Results:  []string{result},
	// 	Filename: filename,
	// }

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderTemplate(w, templates.Error, data400)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		RenderTemplate(w, templates.Error, data400)
		return
	}

	fmt.Println("POST request successful")

	text := strings.TrimSpace(r.FormValue("text"))
	banner := strings.TrimSpace(r.FormValue("banner"))

	if banner == "" {
		banner = "standard"
	}

	text = strings.Replace(text, "\r\n", "\\n", -1)

	if utils.IsAsciiCaractere(text) {
		w.WriteHeader(http.StatusBadRequest)
		RenderTemplate(w, templates.Error, data400)
		return
	}

	var result string

	switch banner {
	case "standard":
		result = src.CreatingArtStandard(text)
	case "shadow":
		result = src.CreatingArtShadow(text, banner)
	default:
		result = src.CreatingArt(text, banner)
	}

	// Enregistrer l'ASCII art dans un fichier
	filename := fmt.Sprintf("%s.txt", uuid.New().String())
	filePath := fmt.Sprintf("ascii-art-files/%s", filename)
	err1 := ioutil.WriteFile(filePath, []byte(result), 0644)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	formData := FormData{
		Text:     text,
		Banner:   banner,
		Results:  []string{result},
		Filename: filename,
	}
	RenderTemplate(w, templates.Index, formData)
}
