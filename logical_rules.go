package validation

// Any returns a Rule that passes when at least one of the given rules passes.
//
// All rules are tried in order; the first pass short-circuits. If every rule fails, ErrValidationAny is returned.
// The inner errors are not propagated.
//
// Fails if:
//   - all supplied rules fail for the value
//
// Examples:
//
//	validation.Any(validation.Email, validation.PhoneE164).Validate("user@example.com") // pass
//	validation.Any(validation.Email, validation.PhoneE164).Validate("+14155552671")      // pass
//	validation.Any(validation.Email, validation.PhoneE164).Validate("notvalid")          // fail
func Any(rules ...Rule) Rule {
	return InputRuleFunc(func(value any, input *InputBag) error {
		for _, r := range rules {
			if err := applyRule(r, value, input); err == nil {
				return nil
			}
		}

		return ErrValidationAny
	})
}

// Not returns a Rule that inverts the result of the given rule.
//
// Passes when the inner rule fails; fails (returning ErrValidationNot) when the inner rule passes.
//
// Fails if:
//   - the wrapped rule passes for the value
//
// Examples:
//
//	validation.Not(validation.Email).Validate("not-an-email") // pass
//	validation.Not(validation.Email).Validate("user@example.com") // fail
//	validation.Not(validation.UUID).Validate("not-a-uuid")    // pass
func Not(r Rule) Rule {
	return InputRuleFunc(func(value any, input *InputBag) error {
		if err := applyRule(r, value, input); err != nil {
			return nil
		}

		return ErrValidationNot
	})
}

// Unless returns an InputRule that applies the given rules only when the condition evaluates to false.
//
// It is the conditional complement of When: rules run when the condition is FALSE.
// The condition language is identical to RequiredIf (comparisons, &&, ||, !, exists(), len()).
// Returns RuleSyntaxError for a malformed condition.
//
// The errors from the inner rules are propagated directly (unlike Any/Not/Each which return sentinels).
//
// Examples:
//
//	validation.Unless(`status == "approved"`, validation.MinLength(10))
//	validation.Unless(`exists(override)`, validation.Required)
func Unless(condition string, rules ...Rule) InputRule {
	return InputRuleFunc(func(value any, input *InputBag) error {
		ok, err := evalCondition(condition, input)
		if err != nil {
			return RuleSyntaxError{Rule: "Unless", Err: err}
		}

		if ok {
			return nil
		}

		for _, r := range rules {
			if err := applyRule(r, value, input); err != nil {
				return err
			}
		}

		return nil
	})
}

// When returns an InputRule that applies the given rules only when the condition evaluates to true.
//
// The condition language is identical to RequiredIf (comparisons, &&, ||, !, exists(), len()).
// Returns RuleSyntaxError for a malformed condition.
//
// The errors from the inner rules are propagated directly (unlike Any/Not/Each which return sentinels).
//
// Examples:
//
//	validation.When(`plan == "paid"`, validation.Regex(`^[A-Z]{2}\d{9}$`), validation.MaxLength(12))
//	validation.When(`country == "US"`, validation.Regex(`^\d{10}$`))
func When(condition string, rules ...Rule) InputRule {
	return InputRuleFunc(func(value any, input *InputBag) error {
		ok, err := evalCondition(condition, input)
		if err != nil {
			return RuleSyntaxError{Rule: "When", Err: err}
		}

		if !ok {
			return nil
		}

		for _, r := range rules {
			if err := applyRule(r, value, input); err != nil {
				return err
			}
		}

		return nil
	})
}

// applyRule dispatches a Rule, routing cross-field rules through ValidateWithInput when an InputBag is available.
// All combinators must call this instead of r.Validate directly so that wrapped InputRules receive the full input.
func applyRule(r Rule, value any, input *InputBag) error {
	if ir, ok := r.(InputRule); ok && input != nil {
		return ir.ValidateWithInput(value, input)
	}

	return r.Validate(value)
}
