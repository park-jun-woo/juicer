//ff:func feature=scan type=extract control=sequence
//ff:what Echo 경로 세그먼트(:param, *wildcard)를 OpenAPI 템플릿으로 변환한다
package echo

import (
	"strings"
)

func echoSegmentToOpenAPI(seg string) string {
	if strings.HasPrefix(seg, ":") {
		return "{" + seg[1:] + "}"
	}
	if strings.HasPrefix(seg, "*") {
		name := seg[1:]
		if name == "" {
			name = "wildcard"
		}
		return "{" + name + "}"
	}
	return seg
}
