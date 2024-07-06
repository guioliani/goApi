package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag.Get("validate")

		rules := strings.Split(tag, ",")

		for _, rule := range rules {
			if rule == "required" {
				if isZero(value) {
					return errors.New(fmt.Sprintf("%s is required", field.Name))
				}
			}

			if strings.HasPrefix(rule, "max=") {
				maxLen := 0
				fmt.Sscanf(rule, "max=%d", &maxLen)
				if str, ok := value.(string); ok {
					if len(str) > maxLen {
						return errors.New(fmt.Sprintf("The %s must have a maximum of 30 characters", field.Name))
					}
				}
			}
		}
	}

	return nil
}

func isZero(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}
