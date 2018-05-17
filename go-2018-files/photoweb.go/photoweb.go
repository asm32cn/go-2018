// photoweb.go
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" " +
			" enctype=\"multipart/form-data\">\n" +
			"\tChoose an image to upload: <input name=\"image\" type=\"file\" />\n" +
			"\t<input type=\"submit\" value=\"Upload\" />\n" +
			"</form>")
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id=" + filename,
			http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	urls := []string{"/upload"}
	io.WriteString(w, "<ul>\n")
	for _, url := range urls {
		io.WriteString(w, "\t<li><a href=\"" + url + "\" />" + url + "</a></li>\n")
	}
	io.WriteString(w, "</ul>")

	io.WriteString(w, "<ul>\n")
	err := filepath.Walk(UPLOAD_DIR, func(path string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() { return nil }
		fileName := f.Name()
		io.WriteString(w, "\t<li><a href=\"/view?id=" + fileName + "\" />" + fileName + "</a></li>\n")
		return nil
	})
	if err != nil {
		println("filepath.Walk() returned %v\n", err)
	}
	io.WriteString(w, "</ul>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}