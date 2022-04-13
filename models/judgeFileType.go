package models

import (
	"mime/multipart"
	"path"

	"github.com/gin-gonic/gin"
)

func JudgeFileType(ctx *gin.Context, file *multipart.FileHeader) {
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true}
	if _, ok := allowExtMap[extName]; !ok {
		ctx.String(200, "file type incorrect.")
		return
	}
}
