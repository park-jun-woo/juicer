//ff:func feature=scan type=extract control=sequence
//ff:what composite-literal body 문자열("Foo{}"/"[]Foo{}")에서 표시 타입명과 base struct명을 분리한다
package echo

import (
	"strings"
)

func compositeLitTypeName(body string) (display string, base string) {
	if !strings.HasSuffix(body, "{}") {
		return "", ""
	}
	name := strings.TrimSuffix(body, "{}")
	display = name
	base = strings.TrimPrefix(name, "[]")
	if strings.ContainsAny(base, ".[]*") || base == "" {
		return "", ""
	}
	return display, base
}
