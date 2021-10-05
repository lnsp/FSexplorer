package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/lnsp/fsexplorer/client"
)

const MaxViewLimit = 1_000_000

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
	static := flag.Bool("static", false, "host static files on root")
	flag.Parse()
	// add endpoint for searching files
	rootDir := os.DirFS(*basedir)
	router := http.NewServeMux()
	router.HandleFunc("/api/view", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		download := r.URL.Query().Get("download") != ""
		log.Printf("viewing with query '%s'", query)
		// check if starts with /, simplify to ./
		if strings.HasPrefix(query, "/") {
			query = strings.Replace(query, "/", "./", 1)
		}
		// list all files in destpath directory
		query = filepath.Clean(query)
		// get file name
		stat, err := fs.Stat(rootDir, query)
		if err != nil {
			http.Error(w, "Could not stat file", http.StatusNotFound)
			return
		}
		if stat.Size() > MaxViewLimit && !download {
			http.Error(w, "File is too large to view", http.StatusInternalServerError)
			return
		}
		file, err := rootDir.Open(query)
		if err != nil {
			http.Error(w, "Could not open file", http.StatusNotFound)
			return
		}
		defer file.Close()
		// write header
		if download {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, stat.Name()))
		}
		w.Header().Set("Content-Length", strconv.Itoa(int(stat.Size())))
		w.WriteHeader(http.StatusOK)
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Could not copy file", http.StatusInternalServerError)
			return
		}
	})
	router.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
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
				files = append(files, fileEntry{
					Name: stat.Name(),
					Path: path,
					Size: stat.Size(),
				})
			}
		}
		// and write as output
		buf, _ := json.Marshal(files)
		w.Write(buf)
	})
	if *static {
		router.Handle("/", http.FileServer(http.FS(client.Dist())))
	}
	if err := http.ListenAndServe(*addr, handlers.CORS()(router)); err != nil {
		log.Fatal(err)
	}
}

func GetFileContentType(fsys fs.FS, path string) (string, error) {
	file, err := fsys.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
