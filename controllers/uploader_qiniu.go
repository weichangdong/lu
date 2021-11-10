package controllers

import (
	"mime/multipart"
	"os"
)

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type QiniuUploader struct {
}

func (u QiniuUploader) upload(file multipart.File, fileHeader *multipart.FileHeader) (url string, err error) {
	return
}
