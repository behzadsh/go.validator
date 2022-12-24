package bag

type ErrorBag map[string][]string

func (b ErrorBag) IsEmpty() bool {
	return len(b) == 0
}

func (b ErrorBag) Add(selector string, msg ...string) {
	b[selector] = append(b[selector], msg...)
}
