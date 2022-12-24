package bag

import (
	"encoding/json"
	"strings"

	"github.com/spf13/cast"
)

type InputBag map[string]any

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

func NewInputBagFromStruct(input any) InputBag {
	b, _ := json.Marshal(input)

	var bag InputBag
	_ = json.Unmarshal(b, &bag)

	return bag
}
