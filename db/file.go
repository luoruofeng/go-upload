package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/luoruofeng/go-upload/db/mysql"
)

const (
	AVAILABLE   = 1
	UNAVAILABLE = 2
	DELETED     = 3
)

func OnFileUploadFinished(filehash string, filename string, filesize int64, location string) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) values (?,?,?,?," + strconv.Itoa(AVAILABLE) + ")")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()
	r, err := stmt.Exec(filehash, filename, filesize, location)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if ra, err := r.RowsAffected(); nil == err {
		if ra <= 0 {
			fmt.Printf("File has been uploaded before. hash:%s", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}
