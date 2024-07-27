package validator

import (
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErros map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErros) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErros == nil {
		v.FieldErros = make(map[string]string)
	}

	if _, exists := v.FieldErros[key]; !exists {
		v.FieldErros[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}
