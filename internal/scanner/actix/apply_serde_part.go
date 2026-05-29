//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what serde 어트리뷰트 한 토큰을 파싱해 결과에 반영한다
package actix

import (
	"strings"
)

func applySerdePart(result *serdeAttr, part string) {
	if part == "default" {
		result.hasDefault = true
	}
	if part == "skip" || part == "skip_deserializing" {
		result.skip = true
	}
	if !strings.HasPrefix(part, "rename") {
		return
	}
	eqIdx := strings.Index(part, "=")
	if eqIdx < 0 {
		return
	}
	val := strings.TrimSpace(part[eqIdx+1:])
	result.rename = unquoteRust(val)
}
