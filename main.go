package main

import (
	"fmt"
	"net/http"

	"github.com/luoruofeng/go-upload/cfg"
	"github.com/luoruofeng/go-upload/handler"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/upload", handler.Upload)
	err := http.ListenAndServe(cfg.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
