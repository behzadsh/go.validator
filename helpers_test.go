package validation

import "errors"

func errorCode(err error) string {
	var ve Error
	if errors.As(err, &ve) {
		return ve.Code()
	}
	return ""
}
