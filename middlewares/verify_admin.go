package middlewares

import (
	"healthSentinel/models"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareAuth(ctx *gin.Context) {
	pathname := strings.Split(ctx.Request.URL.String(), "?")[0]
	nowUsername, err := ctx.Cookie("usrename")
	if err != nil {
		log.Println(err)
		if pathname == "/admin" || pathname == "/admin/" {
			ctx.Redirect(302, "/admin/login")
		}
	}
	nowToken, err := ctx.Cookie("token")
	if err != nil {
		log.Println(err)
		if pathname == "/admin" || pathname == "/admin/" {
			ctx.Redirect(302, "/admin/login")
		}
	}

	if len(nowUsername) > 0 && len(nowToken) > 0 {
		key := []byte("effac6e5")
		strDecrypted, err := models.Decrypt(nowToken, key)
		if err != nil {
			log.Fatal(err)
		}

		str := strings.Split(strDecrypted, "&")
		strLength, _ := strconv.Atoi(str[1])

		if str[0][0:strLength] != nowUsername {
			if pathname == "/admin" || pathname == "/admin/" {
				ctx.Redirect(302, "/admin/login")
			}
		}
	} else {
		if pathname == "/admin" || pathname == "/admin/" {
			ctx.Redirect(302, "/admin/login")
		}
	}
}
