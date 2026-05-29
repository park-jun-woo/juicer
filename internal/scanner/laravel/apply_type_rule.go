//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 타입 규칙(string/integer 등)을 필드에 적용하고 처리 여부를 반환한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyTypeRule(f *scanner.Field, rule string, isNumber *bool) bool {
	t, ok := laravelValidationTypeMap[rule]
	if !ok {
		return false
	}
	f.Type = t
	if t == "integer" || t == "number" {
		*isNumber = true
	}
	return true
}
