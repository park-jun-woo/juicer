//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 포맷 규칙(email/url 등)을 필드에 적용하고 처리 여부를 반환한다
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyFormatRule(f *scanner.Field, rule string) bool {
	if _, ok := laravelValidationFormatMap[rule]; !ok {
		return false
	}
	f.Type = "string"
	f.Validate = appendValidate(f.Validate, rule)
	return true
}
