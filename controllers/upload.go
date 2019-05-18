package controllers

import (
	. "api/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"path"
)

type UploadController struct {
	ApiController
}

var (
	dir      string // 上传文件根路径
	filename string // 上传文件名
)

// Title 上传文件
// Param file
// Param remark
func (c *UploadController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Info("上传文件解析参数错误:", err.Error())
		c.JsonReturn("上传文件解析参数错误:"+err.Error(), "", 500)
	}
	defer f.Close()

	dir = beego.AppConfig.String("StaticUploadPath")
	filename = GetUniqueFileName(h.Filename)
	CheckDirectory(dir + path.Dir(filename))
	err = c.SaveToFile("file", dir+filename)
	if err != nil {
		logs.Info("保存文件错误", err.Error())
		c.JsonReturn("保存文件错误", "", 500)
	}

	result = Result{
		"name": h.Filename,
		"path": filename,
	}
	c.JsonReturn("上传图片成功!", result, 200)
}
