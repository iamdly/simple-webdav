package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func main() {
	log.Println("server start.")

	uName := os.Getenv("USERNAME")
	uPassword := os.Getenv("PASSWORD")

	log.Printf("USERNAME: %s  PASSWORD: %s\n", uName, uPassword)

	fs := &webdav.Handler{
		FileSystem: webdav.Dir("/data"),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != uName || password != uPassword {
			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
			return
		}

		fs.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":6802", nil)
	if err != nil {
		log.Println("Server Error:", err)
	}
}
