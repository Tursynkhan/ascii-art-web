package main

import (
	"fmt"
	"html/template"
	"main/pkg"
	"net/http"
	"os"
)

type art struct {
	Output string
}

func home(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		errorHandler(writer, request, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	if request.URL.Path != "/" {
		errorHandler(writer, request, errStatus{http.StatusNotFound, http.StatusText(http.StatusNotFound)})
		return
	}
	file := "ui/html/index.html"
	ts, err := template.ParseFiles(file)
	if err != nil {
		errorHandler(writer, request, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	err = ts.Execute(writer, nil)
	if err != nil {
		errorHandler(writer, request, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
	}
}

func createAscii(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("inputAscii")

	if r.Method != http.MethodPost {
		errorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	t, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		errorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}

	var s1 art
	banner := r.FormValue("banner")
	txt, err3 := pkg.AsciiPrint(text, banner)
	if err3 == 400 {
		errorHandler(w, r, errStatus{http.StatusBadRequest, http.StatusText(http.StatusBadRequest)})
		return
	}
	if err3 == 500 {
		errorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	s1.Output = txt
	WriteToFile(s1.Output)
	t.Execute(w, s1)
}

func WriteToFile(s string) {
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Unable to create file", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(s)
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	filePath := "text.txt"
	if r.Method != http.MethodGet {
		errorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=Data.txt")
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, filePath)
}
