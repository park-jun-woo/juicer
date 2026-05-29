//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what max:N 규칙을 숫자면 Maximum, 아니면 MaxLength로 적용한다
package laravel

import (
	"strconv"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyMaxRule(f *scanner.Field, value string, isNumber bool) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return
	}
	if isNumber {
		f.Maximum = &n
		return
	}
	f.MaxLength = &n
}
