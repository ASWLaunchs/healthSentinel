package controllers_admin

import (
	"fmt"
	"healthSentinel/models"

	"github.com/ostafen/clover"

	gin "github.com/gin-gonic/gin"
)

type ControllersUserList struct {
}

func (c ControllersUserList) UserList(ctx *gin.Context) {
	queryDate := ctx.Query("queryDate")
	ctx.JSON(200, gin.H{
		"status": true,
		"msg":    getUserList(queryDate),
	})
}

//getUserList() be used calls all members list.
func getUserList(queryDate string) []models.Student {
	var student models.Student
	results := make([]models.Student, 0)

	docs, err := models.DBCloverDB.Query("studentList").Where(clover.Field("createdDate").Eq(queryDate)).FindAll()
	if err != nil {
		panic(err.Error())
	}

	for _, doc := range docs {
		err := doc.Unmarshal(&student)
		if err != nil {
			panic(err.Error())
		}

		results = append(results, student)
	}
	fmt.Println(results)
	return results
}
