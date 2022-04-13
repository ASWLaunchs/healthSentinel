package controllers_office

import (
	"healthSentinel/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ostafen/clover"
)

type ControllersCheck struct {
}

func (c ControllersCheck) Check(ctx *gin.Context) {
	sno := ctx.Query("sno")
	if sno != "" && len(strings.TrimSpace(sno)) > 0 {
		if query(sno) {
			ctx.JSON(200, gin.H{
				"status": true,
				"msg":    "alredy submited today!",
			})
		} else {
			ctx.JSON(200, gin.H{
				"status": false,
				"msg":    "you have no submit.",
			})
		}
	} else {
		ctx.JSON(200, gin.H{
			"status": false,
			"msg":    "formatt error.",
		})
	}

}

func query(sno string) bool {
	// find all completed todos belonging to users with id 5 and 8
	docs, _ := models.DBCloverDB.Query("studentList").Where(clover.Field("sno").Eq(sno)).FindAll()
	var student models.Student

	for _, doc := range docs {
		doc.Unmarshal(&student)
	}
	return student.Completed
}
