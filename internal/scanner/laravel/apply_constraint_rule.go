//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 제약 규칙(max:/min:/in:)을 필드에 적용한다
package laravel

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyConstraintRule(f *scanner.Field, rule string, isNumber bool) {
	if strings.HasPrefix(rule, "max:") {
		applyMaxRule(f, strings.TrimPrefix(rule, "max:"), isNumber)
		return
	}
	if strings.HasPrefix(rule, "min:") {
		applyMinRule(f, strings.TrimPrefix(rule, "min:"), isNumber)
		return
	}
	if strings.HasPrefix(rule, "in:") {
		f.Enum = strings.Split(strings.TrimPrefix(rule, "in:"), ",")
	}
}
