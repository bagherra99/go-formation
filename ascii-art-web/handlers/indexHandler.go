package handlers

import "net/http"


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data404 := struct {
		Title   string
		Message string
	}{
		Title:   "404 - Page Not Found",
		Message: "Erreur de génération de la page",
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		RenderTemplate(w, templates.Error, data404)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderTemplate(w, templates.Error, nil)
		return
	}

	RenderTemplate(w, templates.Index, nil)
}