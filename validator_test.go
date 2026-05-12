package validation

import (
	"errors"
	"sync"
	"testing"
)

func TestSchemaValidate_Map(t *testing.T) {
	schema := New().
		Field("name", Required, MinLength(2)).
		Field("email", Required, Email).
		Field("age", Min[int](18))

	t.Run("valid input", func(t *testing.T) {
		input := map[string]any{
			"name":  "Alice",
			"email": "alice@example.com",
			"age":   25,
		}
		res, err := schema.Validate(input)
		if err != nil {
			t.Fatal(err)
		}
		if res.HasErrors() {
			t.Errorf("expected no errors, got: %v", res.Errors())
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		input := map[string]any{}
		res, err := schema.Validate(input)
		if err != nil {
			t.Fatal(err)
		}
		if !res.HasErrors() {
			t.Fatal("expected errors")
		}
		if len(res.For("name")) == 0 {
			t.Error("expected error for name")
		}
		if len(res.For("email")) == 0 {
			t.Error("expected error for email")
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		input := map[string]any{
			"name":  "Alice",
			"email": "not-an-email",
			"age":   25,
		}
		res, err := schema.Validate(input)
		if err != nil {
			t.Fatal(err)
		}
		emailErrs := res.For("email")
		if len(emailErrs) == 0 {
			t.Error("expected email error")
		}
	})
}

func TestSchemaValidate_Struct(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	schema := New().
		Field("name", Required).
		Field("email", Required, Email)

	t.Run("valid struct", func(t *testing.T) {
		res, err := schema.Validate(User{Name: "Bob", Email: "bob@example.com"})
		if err != nil {
			t.Fatal(err)
		}
		if res.HasErrors() {
			t.Errorf("unexpected errors: %v", res.Errors())
		}
	})

	t.Run("empty struct", func(t *testing.T) {
		res, err := schema.Validate(User{})
		if err != nil {
			t.Fatal(err)
		}
		if !res.HasErrors() {
			t.Fatal("expected errors")
		}
		if len(res.For("name")) == 0 {
			t.Error("expected error for name")
		}
		if len(res.For("email")) == 0 {
			t.Error("expected error for email")
		}
	})

	t.Run("pointer to struct", func(t *testing.T) {
		res, err := schema.Validate(&User{Name: "Carol", Email: "carol@example.com"})
		if err != nil {
			t.Fatal(err)
		}
		if res.HasErrors() {
			t.Errorf("unexpected errors: %v", res.Errors())
		}
	})

	t.Run("nil pointer to struct", func(t *testing.T) {
		var u *User
		res, err := schema.Validate(u)
		if err != nil {
			t.Fatal(err)
		}
		if !res.HasErrors() {
			t.Fatal("expected errors for nil struct pointer")
		}
	})
}

func TestSchemaValidate_MultiError(t *testing.T) {
	schema := New().
		Field("email", Required, Email, MinLength(5))

	res, err := schema.Validate(map[string]any{"email": "x"})
	if err != nil {
		t.Fatal(err)
	}

	emailErrs := res.For("email")
	if len(emailErrs) < 2 {
		t.Errorf("expected at least 2 errors for email, got %d", len(emailErrs))
	}
}

func TestSchemaValidate_NestedPath(t *testing.T) {
	schema := New().
		Field("profile.email", Required, Email)

	t.Run("valid nested", func(t *testing.T) {
		input := map[string]any{
			"profile": map[string]any{"email": "user@example.com"},
		}
		res, err := schema.Validate(input)
		if err != nil {
			t.Fatal(err)
		}
		if res.HasErrors() {
			t.Errorf("unexpected errors: %v", res.Errors())
		}
	})

	t.Run("missing nested", func(t *testing.T) {
		res, err := schema.Validate(map[string]any{})
		if err != nil {
			t.Fatal(err)
		}
		if len(res.For("profile.email")) == 0 {
			t.Error("expected error for profile.email")
		}
	})
}

func TestResult_For(t *testing.T) {
	schema := New().
		Field("a", Required).
		Field("b", Required)

	res, err := schema.Validate(map[string]any{})
	if err != nil {
		t.Fatal(err)
	}

	aErrs := res.For("a")
	bErrs := res.For("b")
	cErrs := res.For("c")

	if len(aErrs) == 0 {
		t.Error("expected error for a")
	}
	if len(bErrs) == 0 {
		t.Error("expected error for b")
	}
	if len(cErrs) != 0 {
		t.Errorf("expected no errors for c, got %d", len(cErrs))
	}
}

func TestFieldError_ErrorsIs(t *testing.T) {
	schema := New().Field("name", Required)
	res, err := schema.Validate(map[string]any{})
	if err != nil {
		t.Fatal(err)
	}

	errs := res.For("name")
	if len(errs) == 0 {
		t.Fatal("expected error")
	}
	fe := errs[0]
	if !errors.Is(fe, ErrValidationRequired) {
		t.Errorf("errors.Is(FieldError, ErrValidationRequired) = false, want true")
	}
}

func TestSchemaValidate_Concurrent(t *testing.T) {
	schema := New().
		Field("name", Required, MinLength(2)).
		Field("email", Required, Email)

	input := map[string]any{
		"name":  "Alice",
		"email": "alice@example.com",
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := schema.Validate(input)
			if err != nil || res.HasErrors() {
				t.Errorf("concurrent Validate failed: err=%v hasErrors=%v", err, res.HasErrors())
			}
		}()
	}
	wg.Wait()
}

func TestSchemaField_Chaining(t *testing.T) {
	s := New()
	s2 := s.Field("a", Required)
	if s != s2 {
		t.Error("Field() should return receiver for chaining")
	}
}
