// Package validate provides a simple way to validate if the input struct
// passes the validation tags. It is used in the controllers of the
// api, but can be used anywhere in the project.
package validate

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Message     string `json:"message"`
}

type ValidationErrors struct {
	Errors []*ErrorResponse `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	var message string
	for _, err := range e.Errors {
		message += fmt.Sprintf("Field %s: %s\n", err.FailedField, err.Message)
	}
	return message
}

// Return an understandabe / UI friendly message that potentially
// can be displayed on the client side, when validation for
// a field fails. Be free to add new tags!
func messageForTag(fe validator.FieldError) string {

	field := fe.Field()
	param := fe.Param()
	tag := fe.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " should be at least " + param + " characters long"
	case "max":
		return field + " should be less than " + param + " characters long"
	}
	return field + " failed validation for " + tag + ", " + param
}

// Struct checks if the validate tags for the input struct pass validation.
func Struct(i interface{}) error {
	err := validator.New().Struct(i)
	if err != nil {
		var errors []*ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.Field(),
				Message:     messageForTag(err),
			})
		}
		return &ValidationErrors{errors}
	}
	return nil
}

// RequestBody returns an error if the body of the response could not be decoded
// or the validate tags for the input interface did not pass validation.
func RequestBody(r *http.Request, body interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		return err
	}
	if err := Struct(body); err != nil {
		return err
	}
	return nil
}
