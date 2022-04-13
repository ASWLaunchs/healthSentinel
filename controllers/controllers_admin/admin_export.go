package controllers_admin

import (
	"healthSentinel/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ControllersExport struct {
}

func (c ControllersExport) Export(ctx *gin.Context) {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	output := "upload/" + models.NowDate()
	archiveName := "static/assets/archive/" + now + ".zip"
	// models.Zip(output, archiveName)

	ctx.HTML(200, "templates_admin/export.html", gin.H{
		"downloadLink": output + archiveName,
	})
}
