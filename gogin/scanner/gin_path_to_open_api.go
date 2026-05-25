//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Gin 경로 패턴(:param)을 OpenAPI 경로 패턴({param})으로 변환한다
package scanner

import (
	"strings"
)

func ginPathToOpenAPI(path string) string {
	segments := strings.Split(path, "/")
	for i, seg := range segments {
		if strings.HasPrefix(seg, ":") {
			segments[i] = "{" + seg[1:] + "}"
		} else if strings.HasPrefix(seg, "*") {
			segments[i] = "{" + seg[1:] + "}"
		}
	}
	return strings.Join(segments, "/")
}

