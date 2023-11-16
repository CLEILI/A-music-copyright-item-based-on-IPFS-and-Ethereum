package main

import (
	"fmt"
	"io"
	"item/ipfs"
	"item/myether"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个Gin路由
	r := gin.Default()
	r.Static("/css", "template/css")
	r.Static("/js", "template/js")
	r.Static("/img", "template/img")
	r.LoadHTMLGlob("template/*.html")
	// 创建一个路由处理器来渲染HTML页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 启动服务器
	r.POST("/tx", func(c *gin.Context) {
		// 使用 PostForm 方法获取表单数据
		file, err := c.FormFile("user-file")
		if err != nil {
			fmt.Println("formfile wrong:", err)
		}
		err = c.SaveUploadedFile(file, "fileupload/"+file.Filename)
		if err != nil {
			fmt.Println("savefile wrong:", err)
		}
		sk := c.PostForm("user-input")
		fmt.Println(sk)
		cid := ipfs.Upload("fileupload/" + file.Filename)
		fmt.Println("cid is :", cid)
		txhash := myether.Storecid(sk, cid)
		fmt.Println("txhash is:", txhash)

		c.HTML(http.StatusOK, "tx.html", gin.H{
			"cidd":    cid,
			"txhashh": txhash,
		})

	})
	r.POST("/te", func(c *gin.Context) {
		cid := c.PostForm("query-input")
		author := myether.Querycid(cid)
		c.HTML(http.StatusOK, "te.html", gin.H{
			"author": author,
		})
	})
	r.POST("/download", func(c *gin.Context) {
		cid := c.PostForm("dinput")
		path := ipfs.Download(cid, "music.mp3")
		file, err := os.Open(path)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Open file error: %s", err.Error()))
			return
		}
		defer file.Close()

		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(path)))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Cache-Control", "no-cache")
		c.Header("Pragma", "no-cache")

		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file error: %s", err.Error()))
		}

		c.HTML(http.StatusOK, "download.html", nil)
	})
	r.Run(":8080")
}
