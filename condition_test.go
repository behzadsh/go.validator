package validation

import (
	"testing"
)

func TestCondTokenize(t *testing.T) {
	tests := []struct {
		input     string
		wantKinds []cTokKind
		wantErr   bool
	}{
		{"", []cTokKind{cTokEOF}, false},
		{"&&", []cTokKind{cTokAND, cTokEOF}, false},
		{"||", []cTokKind{cTokOR, cTokEOF}, false},
		{"==", []cTokKind{cTokEQ, cTokEOF}, false},
		{"!=", []cTokKind{cTokNEQ, cTokEOF}, false},
		{"<=", []cTokKind{cTokLTE, cTokEOF}, false},
		{">=", []cTokKind{cTokGTE, cTokEOF}, false},
		{"<", []cTokKind{cTokLT, cTokEOF}, false},
		{">", []cTokKind{cTokGT, cTokEOF}, false},
		{"!", []cTokKind{cTokNOT, cTokEOF}, false},
		{"(", []cTokKind{cTokLParen, cTokEOF}, false},
		{")", []cTokKind{cTokRParen, cTokEOF}, false},
		{`"hello"`, []cTokKind{cTokString, cTokEOF}, false},
		{`'world'`, []cTokKind{cTokString, cTokEOF}, false},
		{"42", []cTokKind{cTokInt, cTokEOF}, false},
		{"3.14", []cTokKind{cTokFloat, cTokEOF}, false},
		{"true", []cTokKind{cTokBool, cTokEOF}, false},
		{"false", []cTokKind{cTokBool, cTokEOF}, false},
		{"ident", []cTokKind{cTokIdent, cTokEOF}, false},
		{"a.b.c", []cTokKind{cTokIdent, cTokEOF}, false},
		{"role == admin", []cTokKind{cTokIdent, cTokEQ, cTokIdent, cTokEOF}, false},
		{"3.14.15", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			toks, err := condTokenize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("condTokenize(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(toks) != len(tt.wantKinds) {
				t.Fatalf("got %d tokens, want %d: %v", len(toks), len(tt.wantKinds), toks)
			}
			for i, kind := range tt.wantKinds {
				if toks[i].kind != kind {
					t.Errorf("token[%d] kind = %v, want %v", i, toks[i].kind, kind)
				}
			}
		})
	}
}

func TestEvalCondition(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		input     map[string]any
		want      bool
		wantErr   bool
	}{
		{
			name:      "equality true (quoted literal)",
			condition: `role == "admin"`,
			input:     map[string]any{"role": "admin"},
			want:      true,
		},
		{
			name:      "equality false",
			condition: `role == "admin"`,
			input:     map[string]any{"role": "user"},
			want:      false,
		},
		{
			name:      "field vs field comparison true",
			condition: "role == expected",
			input:     map[string]any{"role": "admin", "expected": "admin"},
			want:      true,
		},
		{
			name:      "inequality true",
			condition: `plan != "free"`,
			input:     map[string]any{"plan": "pro"},
			want:      true,
		},
		{
			name:      "inequality false",
			condition: `plan != "free"`,
			input:     map[string]any{"plan": "free"},
			want:      false,
		},
		{
			name:      "AND both true",
			condition: `role == "admin" && plan != "free"`,
			input:     map[string]any{"role": "admin", "plan": "pro"},
			want:      true,
		},
		{
			name:      "AND one false",
			condition: `role == "admin" && plan != "free"`,
			input:     map[string]any{"role": "admin", "plan": "free"},
			want:      false,
		},
		{
			name:      "OR first true",
			condition: `role == "admin" || role == "mod"`,
			input:     map[string]any{"role": "admin"},
			want:      true,
		},
		{
			name:      "OR second true",
			condition: `role == "admin" || role == "mod"`,
			input:     map[string]any{"role": "mod"},
			want:      true,
		},
		{
			name:      "OR both false",
			condition: `role == "admin" || role == "mod"`,
			input:     map[string]any{"role": "user"},
			want:      false,
		},
		{
			name:      "NOT true",
			condition: `!(role == "admin")`,
			input:     map[string]any{"role": "user"},
			want:      true,
		},
		{
			name:      "NOT false",
			condition: `!(role == "admin")`,
			input:     map[string]any{"role": "admin"},
			want:      false,
		},
		{
			name:      "bool literal true",
			condition: "verified == true",
			input:     map[string]any{"verified": true},
			want:      true,
		},
		{
			name:      "bool literal false branch",
			condition: "verified == false",
			input:     map[string]any{"verified": true},
			want:      false,
		},
		{
			name:      "numeric LT true",
			condition: "age < 18",
			input:     map[string]any{"age": 15},
			want:      true,
		},
		{
			name:      "numeric LT false",
			condition: "age < 18",
			input:     map[string]any{"age": 20},
			want:      false,
		},
		{
			name:      "numeric GT true",
			condition: "age > 18",
			input:     map[string]any{"age": 20},
			want:      true,
		},
		{
			name:      "numeric LTE equal",
			condition: "age <= 18",
			input:     map[string]any{"age": 18},
			want:      true,
		},
		{
			name:      "numeric GTE equal",
			condition: "age >= 21",
			input:     map[string]any{"age": 21},
			want:      true,
		},
		{
			name:      "float comparison",
			condition: "score >= 4.5",
			input:     map[string]any{"score": 5.0},
			want:      true,
		},
		{
			name:      "exists true",
			condition: "exists(role)",
			input:     map[string]any{"role": "admin"},
			want:      true,
		},
		{
			name:      "exists false",
			condition: "exists(role)",
			input:     map[string]any{},
			want:      false,
		},
		{
			name:      "dot path comparison true",
			condition: "category.id == 10",
			input:     map[string]any{"category": map[string]any{"id": 10}},
			want:      true,
		},
		{
			name:      "dot path comparison false",
			condition: "category.id == 10",
			input:     map[string]any{"category": map[string]any{"id": 99}},
			want:      false,
		},
		{
			name:      "dot path missing segment",
			condition: `category.name == "books"`,
			input:     map[string]any{},
			want:      false,
		},
		{
			name:      "exists with dot path true",
			condition: "exists(order.status)",
			input:     map[string]any{"order": map[string]any{"status": "pending"}},
			want:      true,
		},
		{
			name:      "exists with dot path false",
			condition: "exists(order.status)",
			input:     map[string]any{"order": map[string]any{}},
			want:      false,
		},
		{
			name:      "len with dot path",
			condition: "len(order.items) > 0",
			input:     map[string]any{"order": map[string]any{"items": []any{"a", "b"}}},
			want:      true,
		},
		{
			name:      "len string GT",
			condition: "len(name) > 2",
			input:     map[string]any{"name": "Alice"},
			want:      true,
		},
		{
			name:      "len slice EQ",
			condition: "len(tags) == 2",
			input:     map[string]any{"tags": []any{"a", "b"}},
			want:      true,
		},
		{
			name:      "len missing field is zero",
			condition: "len(missing) == 0",
			input:     map[string]any{},
			want:      true,
		},
		{
			name:      "grouped expression",
			condition: `(status == "active" || status == "pending") && verified == true`,
			input:     map[string]any{"status": "active", "verified": true},
			want:      true,
		},
		{
			name:      "grouped expression false",
			condition: `(status == "active" || status == "pending") && verified == true`,
			input:     map[string]any{"status": "inactive", "verified": true},
			want:      false,
		},
		{
			name:      "unknown function",
			condition: "unknown(field)",
			input:     map[string]any{},
			wantErr:   true,
		},
		{
			name:      "unclosed paren",
			condition: `(role == "admin"`,
			input:     map[string]any{"role": "admin"},
			wantErr:   true,
		},
		{
			name:      "trailing garbage token",
			condition: `role == "admin" )`,
			input:     map[string]any{"role": "admin"},
			wantErr:   true,
		},
		{
			name:      "non-boolean value without comparison",
			condition: "role",
			input:     map[string]any{"role": "admin"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bag := NewInputBag(tt.input)
			got, err := evalCondition(tt.condition, bag)
			if (err != nil) != tt.wantErr {
				t.Fatalf("evalCondition(%q) error = %v, wantErr %v", tt.condition, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("evalCondition(%q) = %v, want %v", tt.condition, got, tt.want)
			}
		})
	}
}

func TestCondLen(t *testing.T) {
	tests := []struct {
		val  any
		want int
	}{
		{nil, 0},
		{"hello", 5},
		{"", 0},
		{[]any{"a", "b", "c"}, 3},
		{[]int{1, 2}, 2},
		{map[string]any{"a": 1, "b": 2}, 2},
		{42, 0},
	}

	for _, tt := range tests {
		got := condLen(tt.val)
		if got != tt.want {
			t.Errorf("condLen(%v) = %d, want %d", tt.val, got, tt.want)
		}
	}
}

func TestCondToFloat(t *testing.T) {
	tests := []struct {
		val    any
		wantF  float64
		wantOk bool
	}{
		{int(1), 1, true},
		{int8(2), 2, true},
		{int16(3), 3, true},
		{int32(4), 4, true},
		{int64(5), 5, true},
		{uint(6), 6, true},
		{uint8(7), 7, true},
		{uint16(8), 8, true},
		{uint32(9), 9, true},
		{uint64(10), 10, true},
		{float32(1.5), float64(float32(1.5)), true},
		{float64(2.5), 2.5, true},
		{"3.14", 0, false},
		{nil, 0, false},
		{true, 0, false},
	}

	for _, tt := range tests {
		f, ok := condToFloat(tt.val)
		if ok != tt.wantOk {
			t.Errorf("condToFloat(%v) ok = %v, want %v", tt.val, ok, tt.wantOk)
			continue
		}
		if ok && f != tt.wantF {
			t.Errorf("condToFloat(%v) = %v, want %v", tt.val, f, tt.wantF)
		}
	}
}
