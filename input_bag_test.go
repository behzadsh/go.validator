package validation

import (
	"testing"
)

func TestInputBagLookup_Map(t *testing.T) {
	input := map[string]any{
		"name": "Alice",
		"age":  30,
		"profile": map[string]any{
			"email": "alice@example.com",
			"address": map[string]any{
				"city": "Berlin",
			},
		},
		"nilval": nil,
	}
	bag := NewInputBag(input)

	tests := []struct {
		path      string
		wantVal   any
		wantFound bool
	}{
		{"name", "Alice", true},
		{"age", 30, true},
		{"profile.email", "alice@example.com", true},
		{"profile.address.city", "Berlin", true},
		{"nilval", nil, true},
		{"missing", nil, false},
		{"profile.missing", nil, false},
		{"profile.email.deep", nil, false},
		{"", nil, false},
	}

	for _, tt := range tests {
		val, found := bag.Lookup(tt.path)
		if found != tt.wantFound {
			t.Errorf("Lookup(%q) found = %v, want %v", tt.path, found, tt.wantFound)
		}
		if found && val != tt.wantVal {
			t.Errorf("Lookup(%q) val = %v, want %v", tt.path, val, tt.wantVal)
		}
	}
}

func TestInputBagLookup_Struct(t *testing.T) {
	type Address struct {
		City string `json:"city"`
	}
	type Profile struct {
		Email   string  `json:"email"`
		Address Address `json:"address"`
		Hidden  string  `json:"-"`
		NoTag   string
	}
	type User struct {
		Name    string  `json:"name"`
		Profile Profile `json:"profile"`
	}

	u := User{
		Name: "Bob",
		Profile: Profile{
			Email:   "bob@example.com",
			Address: Address{City: "Paris"},
			Hidden:  "secret",
			NoTag:   "visible",
		},
	}

	bag := NewInputBag(u)

	tests := []struct {
		path      string
		wantVal   any
		wantFound bool
	}{
		{"name", "Bob", true},
		{"profile.email", "bob@example.com", true},
		{"profile.address.city", "Paris", true},
		{"profile.NoTag", "visible", true},
		{"profile.Hidden", nil, false}, // json:"-" hides the field
		{"missing", nil, false},
	}

	for _, tt := range tests {
		val, found := bag.Lookup(tt.path)
		if found != tt.wantFound {
			t.Errorf("Lookup(%q) found = %v, want %v", tt.path, found, tt.wantFound)
		}
		if found && val != tt.wantVal {
			t.Errorf("Lookup(%q) val = %v, want %v", tt.path, val, tt.wantVal)
		}
	}
}

func TestInputBagLookup_PointerStruct(t *testing.T) {
	type Inner struct {
		Val string `json:"val"`
	}
	type Outer struct {
		Inner *Inner `json:"inner"`
	}

	t.Run(
		"non-nil pointer", func(t *testing.T) {
			bag := NewInputBag(&Outer{Inner: &Inner{Val: "x"}})
			val, found := bag.Lookup("inner.val")
			if !found || val != "x" {
				t.Errorf("Lookup(inner.val) = %v, %v; want x, true", val, found)
			}
		},
	)

	t.Run(
		"nil outer pointer", func(t *testing.T) {
			var o *Outer
			bag := NewInputBag(o)
			_, found := bag.Lookup("inner.val")
			if found {
				t.Error("Lookup on nil *struct should return false")
			}
		},
	)

	t.Run(
		"nil inner pointer", func(t *testing.T) {
			bag := NewInputBag(&Outer{Inner: nil})
			_, found := bag.Lookup("inner.val")
			if found {
				t.Error("Lookup through nil inner pointer should return false")
			}
		},
	)
}

func TestInputBagLookup_EmbeddedStruct(t *testing.T) {
	type Base struct {
		ID int `json:"id"`
	}
	type Extended struct {
		Base
		Name string `json:"name"`
	}

	bag := NewInputBag(Extended{Base: Base{ID: 42}, Name: "test"})
	val, found := bag.Lookup("id")
	if !found || val != 42 {
		t.Errorf("Lookup(id) on embedded struct = %v, %v; want 42, true", val, found)
	}
}

func TestInputBagLookup_ScalarPointerField(t *testing.T) {
	type Form struct {
		Count *int64   `json:"count"`
		Score *float64 `json:"score"`
	}

	count := int64(42)
	score := 3.14

	t.Run("non-nil pointer returns dereferenced value", func(t *testing.T) {
		bag := NewInputBag(Form{Count: &count, Score: &score})
		val, found := bag.Lookup("count")
		if !found {
			t.Fatal("Lookup(count) should be found")
		}
		if v, ok := val.(int64); !ok || v != 42 {
			t.Errorf("Lookup(count) = %v (%T); want int64(42)", val, val)
		}
		val, found = bag.Lookup("score")
		if !found {
			t.Fatal("Lookup(score) should be found")
		}
		if v, ok := val.(float64); !ok || v != 3.14 {
			t.Errorf("Lookup(score) = %v (%T); want float64(3.14)", val, val)
		}
	})

	t.Run("nil pointer returns not found", func(t *testing.T) {
		bag := NewInputBag(Form{Count: nil, Score: nil})
		_, found := bag.Lookup("count")
		if found {
			t.Error("Lookup(count) on nil *int64 should return false")
		}
		_, found = bag.Lookup("score")
		if found {
			t.Error("Lookup(score) on nil *float64 should return false")
		}
	})
}

func TestInputBagLookup_NonStringKeyedMap(t *testing.T) {
	input := map[int]string{1: "one", 2: "two"}
	bag := NewInputBag(input)
	_, found := bag.Lookup("1")
	if found {
		t.Error("non-string-keyed map should not be traversable")
	}
}
