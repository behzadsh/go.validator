package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type cTokKind int

const (
	cTokEOF cTokKind = iota
	cTokAND
	cTokOR
	cTokNOT
	cTokLParen
	cTokRParen
	cTokEQ
	cTokNEQ
	cTokLT
	cTokGT
	cTokLTE
	cTokGTE
	cTokIdent
	cTokString
	cTokInt
	cTokFloat
	cTokBool
)

type cTok struct {
	kind cTokKind
	val  string
}

//nolint:gocyclo // this function is complex by nature.
func condTokenize(s string) ([]cTok, error) {
	var tokens []cTok
	i := 0
	for i < len(s) {
		if unicode.IsSpace(rune(s[i])) {
			i++
			continue
		}

		switch {
		case strings.HasPrefix(s[i:], "&&"):
			tokens = append(tokens, cTok{cTokAND, "&&"})
			i += 2
		case strings.HasPrefix(s[i:], "||"):
			tokens = append(tokens, cTok{cTokOR, "||"})
			i += 2
		case strings.HasPrefix(s[i:], "=="):
			tokens = append(tokens, cTok{cTokEQ, "=="})
			i += 2
		case strings.HasPrefix(s[i:], "!="):
			tokens = append(tokens, cTok{cTokNEQ, "!="})
			i += 2
		case strings.HasPrefix(s[i:], "<="):
			tokens = append(tokens, cTok{cTokLTE, "<="})
			i += 2
		case strings.HasPrefix(s[i:], ">="):
			tokens = append(tokens, cTok{cTokGTE, ">="})
			i += 2
		case s[i] == '<':
			tokens = append(tokens, cTok{cTokLT, "<"})
			i++
		case s[i] == '>':
			tokens = append(tokens, cTok{cTokGT, ">"})
			i++
		case s[i] == '!':
			tokens = append(tokens, cTok{cTokNOT, "!"})
			i++
		case s[i] == '(':
			tokens = append(tokens, cTok{cTokLParen, "("})
			i++
		case s[i] == ')':
			tokens = append(tokens, cTok{cTokRParen, ")"})
			i++
		case s[i] == '"' || s[i] == '\'':
			quote := s[i]
			i++
			start := i
			for i < len(s) && s[i] != quote {
				i++
			}
			tokens = append(tokens, cTok{cTokString, s[start:i]})
			if i < len(s) {
				i++
			}
		case unicode.IsDigit(rune(s[i])):
			start := i
			isFloat := false
			for i < len(s) && (unicode.IsDigit(rune(s[i])) || s[i] == '.') {
				if s[i] == '.' {
					if isFloat {
						return nil, fmt.Errorf("invalid numeric literal %q", s[start:i+1])
					}
					isFloat = true
				}
				i++
			}
			if isFloat {
				tokens = append(tokens, cTok{cTokFloat, s[start:i]})
			} else {
				tokens = append(tokens, cTok{cTokInt, s[start:i]})
			}
		case unicode.IsLetter(rune(s[i])) || s[i] == '_':
			start := i
			for i < len(s) && (unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) || s[i] == '_' || s[i] == '.') {
				i++
			}
			word := s[start:i]
			switch word {
			case "true", "false":
				tokens = append(tokens, cTok{cTokBool, word})
			default:
				tokens = append(tokens, cTok{cTokIdent, word})
			}
		default:
			i++
		}
	}

	tokens = append(tokens, cTok{cTokEOF, ""})

	return tokens, nil
}

type condParser struct {
	tokens []cTok
	pos    int
	input  *InputBag
}

func (p *condParser) peek() cTok {
	if p.pos < len(p.tokens) {
		return p.tokens[p.pos]
	}

	return cTok{cTokEOF, ""}
}

func (p *condParser) consume() cTok {
	t := p.peek()
	p.pos++

	return t
}

func (p *condParser) parseOr() (bool, error) {
	left, err := p.parseAnd()
	if err != nil {
		return false, err
	}

	for p.peek().kind == cTokOR {
		p.consume()

		right, err := p.parseAnd()
		if err != nil {
			return false, err
		}

		left = left || right
	}

	return left, nil
}

func (p *condParser) parseAnd() (bool, error) {
	left, err := p.parseCmp()
	if err != nil {
		return false, err
	}

	for p.peek().kind == cTokAND {
		p.consume()

		right, err := p.parseCmp()
		if err != nil {
			return false, err
		}

		left = left && right
	}

	return left, nil
}

func (p *condParser) parseCmp() (bool, error) {
	left, err := p.parseUnary()
	if err != nil {
		return false, err
	}

	opTok := p.peek()
	switch opTok.kind {
	case cTokEQ, cTokNEQ, cTokLT, cTokGT, cTokLTE, cTokGTE:
		p.consume()

		right, err := p.parseUnary()
		if err != nil {
			return false, err
		}

		return condCompare(left, opTok.kind, right)
	}

	if b, ok := left.(bool); ok {
		return b, nil
	}

	return false, fmt.Errorf("value %v is not boolean and has no comparison operator", left)
}

func (p *condParser) parseUnary() (any, error) {
	if p.peek().kind == cTokNOT {
		p.consume()

		val, err := p.parseAtom()
		if err != nil {
			return nil, err
		}

		b, ok := val.(bool)
		if !ok {
			return nil, fmt.Errorf("! requires a boolean operand, got %T", val)
		}

		return !b, nil
	}

	return p.parseAtom()
}

func (p *condParser) parseAtom() (any, error) {
	t := p.peek()
	switch t.kind {
	case cTokLParen:
		p.consume()
		b, err := p.parseOr()
		if err != nil {
			return nil, err
		}

		if p.peek().kind != cTokRParen {
			return nil, errors.New("expected closing )")
		}

		p.consume()
		return b, nil
	case cTokIdent:
		if p.pos+1 < len(p.tokens) && p.tokens[p.pos+1].kind == cTokLParen {
			return p.parseCall()
		}

		p.consume()
		val, _ := p.input.Lookup(t.val)
		return val, nil
	case cTokString:
		p.consume()
		return t.val, nil
	case cTokInt:
		p.consume()
		n, _ := strconv.Atoi(t.val) //nolint:errcheck // no need to check we already know t.val is an int
		return n, nil

	case cTokFloat:
		p.consume()
		f, _ := strconv.ParseFloat(t.val, 64) //nolint:errcheck // no need to check we already know t.val is a float
		return f, nil

	case cTokBool:
		p.consume()
		return t.val == "true", nil
	}

	return nil, fmt.Errorf("unexpected token %q", t.val)
}

func (p *condParser) parseCall() (any, error) {
	name := p.consume().val
	p.consume() // consume "("

	if p.peek().kind != cTokIdent {
		return nil, fmt.Errorf("%s() expects a field path argument", name)
	}
	arg := p.consume().val

	if p.peek().kind != cTokRParen {
		return nil, fmt.Errorf("expected ) after argument in %s()", name)
	}

	p.consume()
	switch name {
	case "exists":
		_, ok := p.input.Lookup(arg)
		return ok, nil
	case "len":
		val, ok := p.input.Lookup(arg)
		if !ok {
			return 0, nil
		}
		return condLen(val), nil
	default:
		return nil, fmt.Errorf("unknown function %q", name)
	}
}

func condLen(val any) int {
	if val == nil {
		return 0
	}

	switch v := val.(type) {
	case string:
		return len(v)
	case []any:
		return len(v)
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		return rv.Len()
	}

	return 0
}

func condCompare(left any, op cTokKind, right any) (bool, error) {
	lf, lNum := condToFloat(left)
	rf, rNum := condToFloat(right)
	if lNum && rNum {
		switch op {
		case cTokEQ:
			return lf == rf, nil
		case cTokNEQ:
			return lf != rf, nil
		case cTokLT:
			return lf < rf, nil
		case cTokGT:
			return lf > rf, nil
		case cTokLTE:
			return lf <= rf, nil
		case cTokGTE:
			return lf >= rf, nil
		}
	}

	ls := fmt.Sprintf("%v", left)
	rs := fmt.Sprintf("%v", right)
	switch op {
	case cTokEQ:
		return ls == rs, nil
	case cTokNEQ:
		return ls != rs, nil
	case cTokLT:
		return ls < rs, nil
	case cTokGT:
		return ls > rs, nil
	case cTokLTE:
		return ls <= rs, nil
	case cTokGTE:
		return ls >= rs, nil
	default:
		return false, errors.New("unsupported comparison operator")
	}
}

func condToFloat(v any) (float64, bool) {
	switch n := v.(type) {
	case int:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint8:
		return float64(n), true
	case uint16:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float32:
		return float64(n), true
	case float64:
		return n, true
	default:
		return 0, false
	}
}

func evalCondition(condition string, input *InputBag) (bool, error) {
	tokens, err := condTokenize(condition)
	if err != nil {
		return false, err
	}

	p := &condParser{tokens: tokens, input: input}
	result, err := p.parseOr()
	if err != nil {
		return false, err
	}

	if tok := p.peek(); tok.kind != cTokEOF {
		return false, fmt.Errorf("unexpected token %q", tok.val)
	}

	return result, nil
}
