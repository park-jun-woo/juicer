//ff:func feature=scan type=convert control=selection topic=laravel
//ff:what 플래그 규칙(nullable/required)을 필드에 적용하고 처리 여부를 반환한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyFlagRule(f *scanner.Field, rule string) bool {
	switch rule {
	case "nullable":
		f.Nullable = true
		return true
	case "required":
		f.Validate = appendValidate(f.Validate, "required")
		return true
	}
	return false
}
