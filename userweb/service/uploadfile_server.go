package service

import (
	"crypto/md5"
	"encoding/hex"
	"entry_task/userweb/config"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
)

// FileUpload receive file from user post
func FileUpload(file multipart.File, handler *multipart.FileHeader) string {
	defer file.Close()
	// check file suffix is valid
	if suffix, ok := config.SUFFIX[path.Ext(handler.Filename)]; ok {
		hash := md5.New()
		_, err := io.Copy(hash, file)
		if err != nil {
			log.Printf("hash file load err:%v\n", err)
			return ""
		}
		filePath := fmt.Sprintf("%s/%s%s", config.IMGPATH, hex.EncodeToString(hash.Sum(nil)), suffix)
		_, err = os.Stat(filePath)
		if err == nil {
			return filePath
		}

		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Printf("file creat err:%v\n", err)
			return ""
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Printf("file seek 0,0 err:%v\n", err)
			return ""
		}
		_, err = io.Copy(f, file)
		if err != nil {
			log.Printf("file write err:%v\n", err)
			return ""
		}
		_ = f.Close()
		return filePath
	}
	return ""
}
