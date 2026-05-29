//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what min:N 규칙을 숫자면 Minimum, 아니면 MinLength로 적용한다
package laravel

import (
	"strconv"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyMinRule(f *scanner.Field, value string, isNumber bool) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	if isNumber {
		f.Minimum = &n
		return
	}
	f.MinLength = &n
}
