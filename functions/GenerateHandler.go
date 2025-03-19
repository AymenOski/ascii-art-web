package functions

import (
	"net/http"
	"strconv"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderErrPage(w, http.StatusMethodNotAllowed)
		return
	}

	banner := r.FormValue("banner")
	userText := r.FormValue("text")
	shouldDownload := r.FormValue("download")

	if banner == "" || userText == "" {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}

	if len(userText) > 3000 {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}

	asciiArt, success := GeneratingTheAsciiArt(w, banner, userText)

	if !success {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}

	type result struct {
		Text   string
		Result string
		Banner string
	}
	if shouldDownload != "true" {
		if Exec("result.html", w, result{userText, asciiArt, banner}) != nil {
			RenderErrPage(w, http.StatusInternalServerError)
			return
		}
		return
	}
	// this sets the Content-Type header to text/plain
	w.Header().Set("Content-Type", "text/plain")
	// this tells the browser to download the file.
	w.Header().Set("Content-Disposition", "attachment")
	// this sets the Content-Length header
	w.Header().Set("Content-Length", strconv.Itoa(len(asciiArt)))
	// Write the response body
	w.Write([]byte(asciiArt))
}
