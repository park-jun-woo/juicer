//ff:func feature=scan type=convert control=sequence topic=laravel
//ff:what 단순 영어 단수화(어미 s 제거)를 수행한다
package laravel

import "strings"

// singularize performs naive English singularization (remove trailing 's').
func singularize(s string) string {
	if strings.HasSuffix(s, "ies") {
		return s[:len(s)-3] + "y"
	}
	if strings.HasSuffix(s, "ses") || strings.HasSuffix(s, "xes") || strings.HasSuffix(s, "zes") {
		return s[:len(s)-2]
	}
	if strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") && !strings.HasSuffix(s, "us") {
		return s[:len(s)-1]
	}
	return s
}
