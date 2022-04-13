package controllers_admin

import (
	"github.com/BurntSushi/toml"
	gin "github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ControllersLogin struct {
}

func (c ControllersLogin) Login(ctx *gin.Context) {
	ctx.SetCookie("username", "", 0, "/", "localhost", false, true)
	ctx.SetCookie("token", "", 0, "/", "localhost", false, true)

	bundle := i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("data/i18n/admin/login/active.en.toml")
	bundle.MustLoadMessageFile("data/i18n/admin/login/active.zh-CN.toml")

	lang := ctx.Query("lang")
	accept := ctx.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, lang, accept)

	ctx.HTML(200, "templates_admin/login.html", gin.H{
		"title":       title(localizer),
		"loginTitle":  loginTitle(localizer),
		"textAdmin":   textAdmin(localizer),
		"textPasswd":  textPasswd(localizer),
		"buttonLogin": buttonLogin(localizer),
	})
}

func (c ControllersLogin) LoginFail(ctx *gin.Context) {
	ctx.HTML(200, "templates_admin/loginFail.html", gin.H{})
}

func title(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Title",
			Other: "登录页",
		},
	})
	return result
}

func loginTitle(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TextLoginTitle",
			Other: "管理员登录",
		},
	})
	return result
}

func textAdmin(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TextAdmin",
			Other: "管理员账号",
		},
	})
	return result
}

func textPasswd(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TextPasswd",
			Other: "密码",
		},
	})
	return result
}

func buttonLogin(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ButtonLogin",
			Other: "登录",
		},
	})
	return result
}
