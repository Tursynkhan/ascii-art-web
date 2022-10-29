package delivery

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

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, r, errStatus{http.StatusNotFound, http.StatusText(http.StatusNotFound)})
		return
	}
	file := "./ui/html/index.html"
	ts, err := template.ParseFiles(file)
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
	}
}

func CreateAscii(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("inputAscii")

	if r.Method != http.MethodPost {
		ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}

	var s1 art
	banner := r.FormValue("banner")
	txt, errint := pkg.AsciiPrint(text, banner)
	if errint == 400 {
		ErrorHandler(w, r, errStatus{http.StatusBadRequest, http.StatusText(http.StatusBadRequest)})
		return
	}
	if errint == 500 {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	s1.Output = txt
	WriteToFile(s1.Output)
	t.Execute(w, s1)
	if err != nil {
		ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
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

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filePath := "text.txt"
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=Data.txt")
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, filePath)
}
