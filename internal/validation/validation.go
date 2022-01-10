package validation

import (
	"errors"
	"reflect"
	"strings"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorEn "github.com/go-playground/validator/v10/translations/en"
)

// Error validation
func NewErrValidation() *ErrValidation {
	return &ErrValidation{Errors: make(map[string]string)}
}

type ErrValidation struct {
	Err    error
	Errors map[string]string
}

func (e ErrValidation) Error() string {
	return e.Err.Error()
}

func Validator(s interface{}) error {
	var (
		eng      = english.New()
		uni      = ut.New(eng, eng)
		validate = validator.New()
	)

	trans, found := uni.GetTranslator("en")
	if !found {
		return errors.New("Validator translator not found")
	}

	if err := validatorEn.RegisterDefaultTranslations(validate, trans); err != nil {
		return err
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	customValidation := NewCustomValidation(validate)
	if err := validate.RegisterValidation("mh_email", customValidation.IsEmail); err != nil {
		return nil
	}

	if err := validate.RegisterValidation("mh_login", customValidation.Login); err != nil {
		return nil
	}

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		errValidation := NewErrValidation()
		errValidation.Err = err
		for _, e := range errs {
			errValidation.Errors[e.Field()] = strings.Replace(e.Translate(trans), e.Field(), "", 1)
		}
		return errValidation
	}
	return nil
}

type customValidation struct {
	Validate *validator.Validate
}

func NewCustomValidation(validate *validator.Validate) *customValidation {
	return &customValidation{Validate: validate}
}

func (v *customValidation) IsEmail(fl validator.FieldLevel) bool {
	email := strings.TrimSpace(fl.Field().String())

	if len(email) == 0 {
		return true
	}

	if err := v.Validate.Var(email, "email"); err != nil {
		return false
	}

	return true
}

func (v *customValidation) Login(fl validator.FieldLevel) bool {
	login := strings.TrimSpace(fl.Field().String())

	if len(login) == 0 {
		return true
	}

	if len(login) != 0 && (len(login) < 7 || len(login) > 60) {
		return false
	}

	return true
}
