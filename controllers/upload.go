package controllers

import (
	"api/helpers"
	"github.com/astaxie/beego/logs"
	"time"
)

type UploadController struct {
	ApiController
}

// Title 上传文件
// Param file
// Param remark
func (c *UploadController) Post()  {
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Info("上传文件解析参数错误:", err.Error())
	}
	defer f.Close()

	dir := "static/upload/" + time.Now().Format("2006/01/02/")
	helpers.CheckDirectory(dir)
	filename := dir + h.Filename
	err = c.SaveToFile("file", filename)
	if err != nil {
		logs.Info("保存文件错误", err.Error())
	}

	result := map[string]string{
		"path": filename,
	}
	c.JsonReturn("上传文件接口", result, 200)
}


