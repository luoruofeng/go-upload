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
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.SuccessHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)

	// 用户相关接口
	http.HandleFunc("/", handler.SignInHandler)
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.UserInfoHandler)

	err := http.ListenAndServe(cfg.UploadServiceHost, nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
