package validators

import (
	"errors"

	"github.com/go-playground/validator"
)

type (
	ConcurrentParams struct {
		Type           string `json:"type" query:"type"`
		Items          int    `json:"items" validate:"required" query:"items"`
		ItemsPerWorker int    `json:"items_per_worker" validate:"required" query:"items_per_worker"`
	}

	ConcurrentParamsValidator struct {
		validator *validator.Validate
	}
)

func NewConcurrentParamsValidator(validator *validator.Validate) *ConcurrentParamsValidator {
	return &ConcurrentParamsValidator{validator: validator}
}

func (cv *ConcurrentParamsValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	params := i.(*ConcurrentParams)

	if params.Type != "" && params.Type != "even" && params.Type != "odd" {
		return errors.New(`Type should be one of: ["even", "odd", ""]`)
	}

	return nil
}
