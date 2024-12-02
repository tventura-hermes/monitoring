package validations_helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var StructValidator = validator.New()

func ValidateStruct(v *validator.Validate, s interface{}) bool {
	err := v.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return false
	} else {
		return true
	}
}
