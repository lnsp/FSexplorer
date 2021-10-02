package main

import (
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type fileEntry struct {
	Name      string `json:"name"`
	MIME      string `json:"mime"`
	Size      int64  `json:"size"`
	Directory bool   `json:"isDirectory"`
}

func main() {
	basedir := flag.String("basedir", "/", "base directory to use")
	addr := flag.String("address", "localhost:9876", "address to listen on")
	flag.Parse()
	// setup cors options

	// add endpoint for searching files
	rootDir := os.DirFS(*basedir)
	router := mux.NewRouter()
	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// go through filesystem
		query := r.URL.Query().Get("q")
		log.Printf("searching with query '%s'", query)
		// check if starts with /, simplify to ./
		if strings.HasPrefix(query, "/") {
			query = strings.Replace(query, "/", "./", 1)
		}
		// list all files in destpath directory
		query = filepath.Clean(query)
		entries, err := fs.ReadDir(rootDir, query)
		if err != nil {
			log.Println("got error while reading dir:", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		files := make([]fileEntry, 0, len(entries))
		for _, e := range entries {
			if e.IsDir() {
				files = append(files, fileEntry{
					Name:      e.Name(),
					Directory: true,
				})
			} else {
				extension := filepath.Ext(e.Name())
				mimeType := mime.TypeByExtension(extension)
				stats, err := fs.Stat(rootDir, filepath.Join(query, e.Name()))
				if err != nil {
					log.Fatal(err)
				}
				files = append(files, fileEntry{
					Name: e.Name(),
					MIME: mimeType,
					Size: stats.Size(),
				})
			}
		}
		// and write as output
		buf, _ := json.Marshal(files)
		w.Write(buf)
	})
	if err := http.ListenAndServe(*addr, handlers.CORS()(router)); err != nil {
		log.Fatal(err)
	}
}
