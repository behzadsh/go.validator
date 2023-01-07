package bag

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cast"
)

// InputBag is a custom type representing the validation input.
type InputBag map[string]any

// Get returns the associated value to the given selector. The selector could
// have dot in its value for selecting nested values. E.g. `user.settings.avatar`.
func (b InputBag) Get(selector string) (any, bool) {
	parts := strings.Split(selector, ".")

	if len(parts) == 1 {
		v, ok := b[selector]
		return v, ok
	}

	base := b[parts[0]]
	for i := 1; i < len(parts); i++ {
		if k, err := cast.ToIntE(parts[i]); err == nil {
			tmp, err := cast.ToSliceE(base)
			if err != nil || len(tmp) <= k {
				return nil, false
			}

			base = tmp[k]
			continue
		}

		tmp, ok := base.(map[string]any)
		if !ok {
			return nil, false
		}

		v, ok := tmp[parts[i]]
		if !ok {
			return nil, false
		}
		base = v
	}

	return base, base != nil
}

// Has checks if the given selector is exists in the input bag. The selector
// can contain dots in its value for nested values.
func (b InputBag) Has(selector string) bool {
	parts := strings.Split(selector, ".")

	if len(parts) == 1 {
		_, ok := b[selector]
		return ok
	}

	base := b[parts[0]]
	for i := 1; i < len(parts); i++ {
		if k, err := cast.ToIntE(parts[i]); err == nil {
			tmp, err := cast.ToSliceE(base)
			if err != nil || len(tmp) <= k {
				return false
			}

			base = tmp[k]
			continue
		}

		tmp, ok := base.(map[string]any)
		if !ok {
			return false
		}

		v, ok := tmp[parts[i]]
		if !ok {
			return false
		}
		base = v
	}

	return base != nil
}

// NewInputBagFromStruct converts the given input struct into InputBag.
// Note that since we use json.Marshal and json.Unmarshal for this conversion
// the given struct must have exported field, and if the exported fields have
// json tag, keep in mind that the InputBag keys are the same as the tags.
func NewInputBagFromStruct(input any) InputBag {
	b, _ := json.Marshal(input)

	var bag InputBag
	_ = json.Unmarshal(b, &bag)

	return bag
}
