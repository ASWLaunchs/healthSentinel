package controllers_office

import (
	"healthSentinel/models"
	"net/http"
	"path"
	"sync"
	"time"

	gin "github.com/gin-gonic/gin"
	clover "github.com/ostafen/clover"
)

type ControllersUpload struct {
}

var mutex sync.Mutex

func (c ControllersUpload) Upload(ctx *gin.Context) {
	studentID := ctx.PostForm("studentID")
	healthCode, err1 := ctx.FormFile("healthCode")
	pathCode, err2 := ctx.FormFile("pathCode")
	closeCode, err3 := ctx.FormFile("closeCode")

	//create dir.
	var folderPath string = "upload/" + models.NowDate() + "/" + studentID
	models.CreateMutiDir(folderPath)

	var dst1, dst2, dst3 string
	// storage three images.
	if err1 == nil {
		dst1 = path.Join(folderPath, healthCode.Filename)
		ctx.SaveUploadedFile(healthCode, dst1)
	}
	if err2 == nil {
		dst2 = path.Join(folderPath, pathCode.Filename)
		ctx.SaveUploadedFile(pathCode, dst2)
	}
	if err3 == nil {
		dst3 = path.Join(folderPath, closeCode.Filename)
		ctx.SaveUploadedFile(closeCode, dst3)
	}

	go add(studentID, dst1, dst2, dst3)
	time.Sleep(time.Second)

	ctx.HTML(http.StatusOK, "templates_office/feedback.html", gin.H{
		"feedbackMsg": "提交成功！", "studentID": studentID})
}

func add(sno string, healthCodePath string, pathCodePath string, closeCodePath string) {
	mutex.Lock()
	doc := clover.NewDocument()

	doc.Set("sno", sno)
	doc.Set("createdDate", models.NowDate())
	doc.Set("completed", true)
	doc.Set("healthCodePath", healthCodePath)
	doc.Set("pathCodePath", pathCodePath)
	doc.Set("closeCodePath", closeCodePath)

	models.DBCloverDB.InsertOne("studentList", doc)
	mutex.Unlock()
}
