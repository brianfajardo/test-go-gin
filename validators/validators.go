package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

func ValidateProfanity(field validator.FieldLevel) bool {
	return !strings.Contains(field.Field().String(), "fuck")
}
