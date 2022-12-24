package bag

type ErrorBag map[string][]string

func (b ErrorBag) IsEmpty() bool {
	return len(b) == 0
}

func (b ErrorBag) Add(selector string, msg ...string) {
	b[selector] = append(b[selector], msg...)
}

func (b ErrorBag) FirstOf(name string) string {
	msg, ok := b[name]
	if !ok {
		return ""
	}

	return msg[0]
}

func (b ErrorBag) Has(name string) bool {
	_, ok := b[name]

	return ok
}
