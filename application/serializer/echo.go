package serializer

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type (
	CustomValidator struct {
		Validator  *validator.Validate
		Translator ut.Translator
	}

	ErrorResponse struct {
		Message interface{} `json:"message"`
	}

	SuccessResponse struct {
		Data interface{} `json:"data"`
	}

	ListResponse struct {
		Results interface{} `json:"results"`
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
