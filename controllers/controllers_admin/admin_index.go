package controllers_admin

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	gin "github.com/gin-gonic/gin"
)

type ControllersIndex struct {
}

func (c ControllersIndex) Index(ctx *gin.Context) { //i18n
	bundle := i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("data/i18n/admin/index/active.en.toml")
	bundle.MustLoadMessageFile("data/i18n/admin/index/active.zh-CN.toml")

	lang := ctx.Query("lang")
	accept := ctx.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, lang, accept)

	ctx.HTML(200, "templates_admin/index.html", gin.H{
		"slogan":            slogan(localizer),
		"buttonExport":      buttonExport(localizer),
		"historyData":       historyData(localizer),
		"buttonExit":        buttonExit(localizer),
		"tableSno":          tableSno(localizer),
		"tableUsername":     tableUsername(localizer),
		"tableSubmitedTime": tableSubmitedTime(localizer),
		"tableStatus":       tableStatus(localizer),
	})
}

//head -> slogan
func slogan(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Slogan",
			Other: "健康守卫24h无休收集观察人员防疫数据，三码一核酸，守护你我他",
		},
	})
	return result
}

func buttonExport(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ButtonExport",
			Other: "健康守卫24h无休收集观察人员防疫数据，三码一核酸，守护你我他",
		},
	})
	return result
}

func historyData(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HistoryData",
			Other: "历史数据",
		},
	})
	return result
}

func buttonExit(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ButtonExit",
			Other: "注销",
		},
	})
	return result
}

func tableSno(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TableSno",
			Other: "学号",
		},
	})
	return result
}

func tableUsername(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TableUsername",
			Other: "姓名",
		},
	})
	return result
}

func tableSubmitedTime(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TableSubmitedTime",
			Other: "提交时间",
		},
	})
	return result
}

func tableStatus(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "TableStatus",
			Other: "提交状态",
		},
	})
	return result
}
