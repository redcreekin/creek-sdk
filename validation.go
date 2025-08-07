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

var projectTypes = map[string]bool{
	"build":    true,
	"deploy":   true,
	"data":     true,
	"test":     true,
	"sales":    true,
	"security": true,
}

func initValidator() {
	validate = validator.New()
	// Register idValidPrefix validation
	validate.RegisterValidation("idValidPrefix", idValidation)
	// Register projectType validation
	validate.RegisterValidation("projectType", projectTypeValidation)
	// Register semanticVersion validation
	validate.RegisterValidation("semanticVersion", semanticVersionValidation)
	// Register username validation
	validate.RegisterValidation("username", usernameValidation)
	trans, _ := uni.GetTranslator("en")

	// Register password validation
	validate.RegisterValidation("user_password", passwordValidation)

	// Register the translator for the validator
	enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("idValidPrefix", trans, idRegisterTranslation, idTranslationFunc)
	validate.RegisterTranslation("projectType", trans, projectTypeRegisterTranslation, projectTypeTranslationFunc)
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

func semanticVersionValidation(fl validator.FieldLevel) bool {
	// Example pattern: 1.0.0, 1.0.0-alpha, 1.0.0+build.123
	pattern := `^(\d+\.\d+\.\d+(-[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*)?(\+[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*)?)$`
	versionValue := fl.Field().String()
	if versionValue == "" {
		return false
	}
	match, _ := regexp.MatchString(pattern, versionValue)
	return match
}

func usernameValidation(fl validator.FieldLevel) bool {
	// Example pattern: alphanumeric characters, underscores, and hyphens
	pattern := `^[a-zA-Z0-9_]{3,20}$`
	usernameValue := fl.Field().String()
	if usernameValue == "" {
		return false
	}
	match, _ := regexp.MatchString(pattern, usernameValue)
	return match
}

func passwordValidation(fl validator.FieldLevel) bool {
	// Example pattern: at least 8 characters, one uppercase, one lowercase, one digit, and one special character
	pattern := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	passwordValue := fl.Field().String()
	if passwordValue == "" {
		return false
	}
	match, _ := regexp.MatchString(pattern, passwordValue)
	return match
}

func projectTypeValidation(fl validator.FieldLevel) bool {
	projectType := fl.Field().String()
	return projectTypes[projectType]
}

func idRegisterTranslation(utrans ut.Translator) error {
	return utrans.Add("idValidPrefix", "{0} must start with {1}-<number> and end with a number", true)
}

func idTranslationFunc(utrans ut.Translator, fe validator.FieldError) string {
	t, _ := utrans.T("idValidPrefix", fe.Field(), fe.Param())
	return t
}

func projectTypeRegisterTranslation(utrans ut.Translator) error {
	return utrans.Add("projectType", "{0} must be one of the following: build, deploy, data, test, sales, security", true)
}

func projectTypeTranslationFunc(utrans ut.Translator, fe validator.FieldError) string {
	t, _ := utrans.T("projectType", fe.Field())
	return t
}
