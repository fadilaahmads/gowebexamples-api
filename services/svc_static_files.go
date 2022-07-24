package services

import "net/http"

func StaticFiles() *http.Handler {
	fs := http.FileServer(http.Dir("assets/"))
	return &fs
}
