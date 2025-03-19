package functions

import (
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderErrPage(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		RenderErrPage(w, http.StatusNotFound)
		return
	}

	if Exec("index.html", w, nil) != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
}
