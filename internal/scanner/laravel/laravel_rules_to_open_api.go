//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Laravel 유효성 규칙 문자열 목록을 scanner.Field로 변환한다
package laravel

import (
	"strconv"
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// laravelRulesToField converts a field name and its Laravel validation rules into a scanner.Field.
func laravelRulesToField(name string, rules []string) scanner.Field {
	f := scanner.Field{
		Name: name,
		JSON: name,
		Type: "string", // default
	}

	isString := false
	isNumber := false

	for _, raw := range rules {
		rule := strings.TrimSpace(raw)

		// Check for type rules
		if t, ok := laravelValidationTypeMap[rule]; ok {
			f.Type = t
			if t == "string" {
				isString = true
			}
			if t == "integer" || t == "number" {
				isNumber = true
			}
			continue
		}

		// Check for format rules
		if fmt, ok := laravelValidationFormatMap[rule]; ok {
			f.Type = "string"
			isString = true
			f.Validate = appendValidate(f.Validate, rule)
			_ = fmt // format stored via Validate for now
			continue
		}

		// nullable
		if rule == "nullable" {
			f.Nullable = true
			continue
		}

		// required — handled at caller level (required array)
		if rule == "required" {
			f.Validate = appendValidate(f.Validate, "required")
			continue
		}

		// max:N
		if strings.HasPrefix(rule, "max:") {
			n, err := strconv.Atoi(strings.TrimPrefix(rule, "max:"))
			if err == nil {
				if isNumber {
					f.Maximum = &n
				} else {
					f.MaxLength = &n
				}
			}
			continue
		}

		// min:N
		if strings.HasPrefix(rule, "min:") {
			n, err := strconv.Atoi(strings.TrimPrefix(rule, "min:"))
			if err == nil {
				if isNumber {
					f.Minimum = &n
				} else {
					f.MinLength = &n
				}
			}
			continue
		}

		// in:a,b,c
		if strings.HasPrefix(rule, "in:") {
			values := strings.Split(strings.TrimPrefix(rule, "in:"), ",")
			f.Enum = values
			continue
		}
	}

	_ = isString
	return f
}

// appendValidate appends a validation rule to the existing validate string.
func appendValidate(existing, rule string) string {
	if existing == "" {
		return rule
	}
	return existing + "," + rule
}
