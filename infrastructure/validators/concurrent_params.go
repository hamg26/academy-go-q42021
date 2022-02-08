package validators

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator"
)

type (
	ConcurrentParams struct {
		Type           string `json:"type" query:"type"`
		Items          int    `json:"items" validate:"numeric" query:"items"`
		ItemsPerWorker int    `json:"items_per_worker" validate:"numeric" query:"items_per_worker"`
	}

	ConcurrentParamsValidator struct {
		validator *validator.Validate
	}
)

// Creates a new Validator for the params of the concurrent endpoint
func NewConcurrentParamsValidator(validator *validator.Validate) *ConcurrentParamsValidator {
	return &ConcurrentParamsValidator{validator: validator}
}

// Validates the params needed for the concurrent endpoint
func (cv *ConcurrentParamsValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	params := i.(*ConcurrentParams)

	if params.Type != "" && params.Type != "even" && params.Type != "odd" {
		return errors.New(`Type should be one of: ["even", "odd", ""]`)
	}

	// Maybe thi is a bit intrincate just to validate a couple of params
	// But I wanted to try this kind of dyamic access of a structure properties
	checkItemsParam := func(paramName string) error {
		r := reflect.ValueOf(params)
		v := int(reflect.Indirect(r).FieldByName(paramName).Int())
		if v != -1 && v < 1 {
			return errors.New(paramName + ` should be -1 or an integer greater than 0`)
		}
		return nil
	}

	for _, paramName := range []string{"Items", "ItemsPerWorker"} {
		if err := checkItemsParam(paramName); err != nil {
			return err
		}
	}

	return nil
}
