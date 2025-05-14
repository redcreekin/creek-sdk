package sdk

import (
	"fmt"
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var enLocale = en.New()
var uni = ut.New(enLocale, enLocale)

func initValidator() {
	validate = validator.New()
	// Register idValidPrefix validation
	validate.RegisterValidation("idValidPrefix", idValidation)
	trans, _ := uni.GetTranslator("en")

	// Register the translator for the validator
	enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("idValidPrefix", trans, idRegisterTranslation, idTranslationFunc)
}

func GetValidator() *validator.Validate {
	if validate == nil {
		initValidator()
	}
	return validate
}

func idValidation(fl validator.FieldLevel) bool {
	param := fl.Param()
	idValue := fl.Field().String()
	if idValue == "" {
		return false // Allow empty string
	}
	pattern := fmt.Sprintf("^%s-\\d+$", param)
	match, _ := regexp.MatchString(pattern, idValue)
	return match
}

func idRegisterTranslation(utrans ut.Translator) error {
	return utrans.Add("idValidPrefix", "{0} must start with {1}-<number> and end with a number", true)
}

func idTranslationFunc(utrans ut.Translator, fe validator.FieldError) string {
	t, _ := utrans.T("idValidPrefix", fe.Field(), fe.Param())
	return t
}
