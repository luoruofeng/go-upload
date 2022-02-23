package meta

import (
	db "github.com/luoruofeng/go-upload/db"
)

type FileMeta struct {
	FileSha1   string
	FileName   string
	CreateTime string
	FileSize   int64
	Location   string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFile(fm FileMeta) {
	fileMetas[fm.FileSha1] = fm
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//insert and update from db
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}
