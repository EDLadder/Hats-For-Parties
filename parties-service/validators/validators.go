package validators

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func ValidateInputs(dataSet interface{}) (bool, map[string][]string) {
	validate = validator.New()
	err := validate.Struct(dataSet)

	if err != nil {
		errors := make(map[string][]string)
		if err, ok := err.(*validator.InvalidValidationError); ok {
			errors["panic"] = append(errors["panic"], err.Error())
			return false, errors
		}
		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {

			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "alpha":
				errors[name] = append(errors[name], "The "+name+" should contain only letters")
				break
			case "numeric":
				errors[name] = append(errors[name], "The "+name+" should be number")
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
		}

		return false, errors
	}
	return true, nil
}
