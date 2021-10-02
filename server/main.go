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
	Path      string `json:"path"`
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
	router.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		log.Printf("viewing with query '%s'", query)
		// check if starts with /, simplify to ./
		if strings.HasPrefix(query, "/") {
			query = strings.Replace(query, "/", "./", 1)
		}
		// list all files in destpath directory
		query = filepath.Clean(query)
		content, _ := fs.ReadFile(rootDir, query)
		w.Write(content)
	})
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
		paths := make([]string, len(entries))
		for i, e := range entries {
			paths[i] = filepath.Join(query, e.Name())
		}
		files := make([]fileEntry, 0, len(entries))
		for _, path := range paths {
			stat, err := fs.Stat(rootDir, path)
			if err != nil {
				log.Printf("failed to stat file '%s': %v", path, err)
				continue
			}
			if stat.IsDir() {
				files = append(files, fileEntry{
					Name:      stat.Name(),
					Path:      path,
					Directory: true,
				})
			} else {
				extension := filepath.Ext(stat.Name())
				mimeType := mime.TypeByExtension(extension)
				files = append(files, fileEntry{
					Name: stat.Name(),
					Path: path,
					MIME: mimeType,
					Size: stat.Size(),
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
