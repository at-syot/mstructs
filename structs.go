package structs

import (
	"reflect"
)

func ToMap(o interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	kind := v.Kind()

	if kind == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			fn := t.Field(i).Name
			fv := v.Field(i)

			result[fn] = fv
		}
	}

	return result
}
