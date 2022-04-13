package controllers_verify

import (
	"fmt"
	"healthSentinel/models"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/ini.v1"

	"github.com/gin-gonic/gin"
)

type ControllersVerify struct {
}

func (c ControllersVerify) Admin(ctx *gin.Context) {
	nowUsername := ctx.PostForm("username")
	nowPasswd := ctx.PostForm("passwd")

	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	address := cfg.Section("website").Key("address").String()
	adminUsername := cfg.Section("simpleAuth").Key("username").String()
	adminPasswd := cfg.Section("simpleAuth").Key("passwd").String()

	if nowUsername == adminUsername && nowPasswd == adminPasswd {
		key := []byte("effac6e5")
		str := nowUsername + strconv.FormatInt(time.Now().UnixNano(), 10) + "&" + strconv.Itoa(len(nowUsername))
		strEncrypted, err := models.Encrypt(str, key)
		if err != nil {
			log.Fatal(err)
		}
		ctx.SetCookie("username", nowUsername, 3600, "/", address, false, true)
		ctx.SetCookie("token", strEncrypted, 3600, "/", address, false, true)
		ctx.Redirect(302, "/admin?"+models.NowDate())
	} else {
		ctx.Redirect(302, "/admin/loginFail")
	}
}
