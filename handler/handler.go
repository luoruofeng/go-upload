package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	meta "github.com/luoruofeng/go-upload/meta"
	util "github.com/luoruofeng/go-upload/util"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fd, fh, err := r.FormFile("file")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer fd.Close()

		fileMeta := meta.FileMeta{
			FileName:   fh.Filename,
			FileSize:   fh.Size,
			Location:   "/tmp/" + fh.Filename,
			CreateTime: time.Now().Format("2006-1-2 15:04:05"),
		}

		des, err := os.Create(fileMeta.Location)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer des.Close()

		r, err := io.Copy(des, fd)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		des.Seek(0, 0)
		fileSha1 := util.FileSha1(des)
		fileMeta.FileSha1 = fileSha1
		des.Seek(0, 0)

		meta.Update(fileMeta)

		fmt.Printf("save %d size\n", r)
		w.Write([]byte("save sucess!"))

		// if suc {
		// 	http.Redirect(w, r, "/static/view/home.html", http.StatusFound)
		// } else {
		// 	w.Write([]byte("Upload Failed."))
		// }

	} else if r.Method == http.MethodGet {
		c, err := ioutil.ReadFile("./static/vm/upload.html")
		if err != nil {
			io.WriteString(w, err.Error())
		}
		io.WriteString(w, string(c))
	}
}

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Success!!")
}

// sha1sum /tmp/filename
// 127.0.0.1:8080/file/meta?filehash=xxxx
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	// filehash := r.Form["filehash"][0]
	fmeta := meta.GetFileMeta(filehash)
	if fmeta != (meta.FileMeta{}) {
		data, err := json.Marshal(fmeta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	} else {
		w.Write([]byte(`{"code":-1,"msg":"no such file!"}`))
	}
}
