package controllers_office

import (
	"fmt"

	"github.com/BurntSushi/toml"
	gin "github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ControllersIndex struct {
}

func (c ControllersIndex) Index(ctx *gin.Context) {
	//i18n
	bundle := i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("data/i18n/office/active.en.toml")
	bundle.MustLoadMessageFile("data/i18n/office/active.zh-CN.toml")

	lang := ctx.Query("lang")
	accept := ctx.GetHeader("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, lang, accept)

	fmt.Println(lang, accept)

	ctx.HTML(200, "templates_office/index.html", gin.H{
		"title":        title(localizer),
		"contentTitle": conetentTitle(localizer),
		"sno":          sno(localizer),
		"healthCode":   healthCode(localizer),
		"pathCode":     pathCode(localizer),
		"closeCode":    closeCode(localizer),
		"button":       button(localizer),
		"checkRes":     checkRes(localizer),
	})
}

//head -> title
func title(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Title",
			Other: "健康哨兵",
		},
	})
	return result
}

//body -> h2 conetentTitle
func conetentTitle(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "ContentTitle",
			Other: "2022届18计本三码收集",
		},
	})
	return result
}

//body -> sno
func sno(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Sno",
			Other: "学号",
		},
	})
	return result
}

//body-> healthCode
func healthCode(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "HealthCode",
			Other: "健康码",
		},
	})
	return result
}

//body-> pathCode
func pathCode(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PathCode",
			Other: "行程码",
		},
	})
	return result
}

//body-> closeCode
func closeCode(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "CloseCode",
			Other: "密接码",
		},
	})
	return result
}

//body-> button
func button(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "Button",
			Other: "提交",
		},
	})
	return result
}

func checkRes(localizer *i18n.Localizer) string {
	result := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "CheckRes",
			Other: "您今天已经提交过三码了！",
		},
	})
	return result
}
