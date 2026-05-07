package rules

import "reflect"

func indirectValue(a any) reflect.Value {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Pointer {
		return v
	}

	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}

	return v
}
