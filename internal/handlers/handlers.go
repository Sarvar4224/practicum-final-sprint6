package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "no file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "file read error", http.StatusInternalServerError)
		return
	}
	result, err := service.ConvertMorse(string(data))
	if err != nil {
		http.Error(w, "convert error", http.StatusInternalServerError)
		return
	}

	newName := time.Now().UTC().Format("20060102") + filepath.Ext(handler.Filename)
	out, err := os.Create(newName)
	if err != nil {
		http.Error(w, "can't create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := out.Write([]byte(result)); err != nil {
		http.Error(w, "file write error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
