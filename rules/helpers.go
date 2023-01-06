package rules

import "reflect"

func indirectValue(a any) reflect.Value {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		return v
	}

	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return v
}
