package controllers

import (
	"io/ioutil"
	"net/http"
	"os"

	"wblog/helpers"
	"wblog/system"

	"github.com/gin-gonic/gin"
)

func BackupPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	err = Backup()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func RestorePost(c *gin.Context) {
	var (
		fileName  string
		fileUrl   string
		err       error
		res       = gin.H{}
		resp      *http.Response
		bodyBytes []byte
	)
	defer writeJSON(c, res)
	fileName = c.PostForm("fileName")
	if fileName == "" {
		res["message"] = "fileName cannot be empty."
		return
	}
	fileUrl = system.GetConfiguration().QiniuFileServer + fileName
	resp, err = http.Get(fileUrl)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	bodyBytes, err = helpers.Decrypt(bodyBytes, system.GetConfiguration().BackupKey)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = ioutil.WriteFile(fileName, bodyBytes, os.ModePerm)
	if err == nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func Backup() (err error) {
	return
}
