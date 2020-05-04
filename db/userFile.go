package db

import(
	mydb "fileSystem/db/mysql"
	"fmt"
	"time"
)

//用户表结构体
type UserFile struct {
	Username	string
	FileHash	string
	FileName 	string
	FileSzie 	int64
	UploadAt	string
	LastUpdated	string
}

//更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file(`user_name`,`file_sha1`,`file_name`,`file_size`,`upload_at`)values (?,?,?,?,?)")
	if err != nil{
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil{
		fmt.Println("1", err)
		return false
	}
	fmt.Println("SUCCESS!")
	return true
}

//批量获取用户文件信息
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	stmt, err := mydb.DBConn().Prepare(
			"select file_sha1,file_name,file_size,upload_at," +
			"last_update from tbl_user_file where user_name=? limit ?")
	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(username, limit)
	if err != nil{
		return nil, err
	}

	var userFiles []UserFile
	for(rows.Next()){
		ufile := UserFile{}
		rows.Scan(&ufile.FileHash, &ufile.FileName, &ufile.FileSzie, &ufile.UploadAt, &ufile.LastUpdated)
		if err != nil{
			fmt.Println(err.Error())
			break
		}
		userFiles = append(userFiles, ufile)
	}
	return userFiles, nil
}