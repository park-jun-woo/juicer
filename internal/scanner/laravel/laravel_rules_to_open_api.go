//ff:func feature=scan type=convert control=iteration dimension=1 topic=laravel
//ff:what Laravel 유효성 규칙 문자열 목록을 scanner.Field로 변환한다
package laravel

import (
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
	isNumber := false
	for _, raw := range rules {
		applyValidationRule(&f, strings.TrimSpace(raw), &isNumber)
	}
	return f
}
