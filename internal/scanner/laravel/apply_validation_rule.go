//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 단일 Laravel 유효성 규칙을 필드에 적용한다(타입/포맷/제약/enum)
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyValidationRule(f *scanner.Field, rule string, isNumber *bool) {
	if applyTypeRule(f, rule, isNumber) {
		return
	}
	if applyFormatRule(f, rule) {
		return
	}
	if applyFlagRule(f, rule) {
		return
	}
	applyConstraintRule(f, rule, *isNumber)
}
