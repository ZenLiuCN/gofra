package http

import (
	"net/http"
	"os"
	"path/filepath"
)

// Spa [2]string{wwwRoot,indexFile}
type Spa [2]string

func (h Spa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h[0], r.URL.Path)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h[0], h[1])) // when not found or is dir,redirect to index
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h[0])).ServeHTTP(w, r)
}
