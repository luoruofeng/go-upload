package meta

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

func Update(fm FileMeta) {
	fileMetas[fm.FileSha1] = fm
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
