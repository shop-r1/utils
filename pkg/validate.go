package pkg

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitValidate() {
	en1 := en.New()
	zh1 := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(zh1, en1, zh_tw)
	Validate = validator.New()
}
