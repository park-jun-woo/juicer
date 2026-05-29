//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Echo 경로 패턴(:param)을 OpenAPI 경로 패턴({param})으로 변환한다
package echo

import (
	"strings"
)

func echoPathToOpenAPI(path string) string {
	segments := strings.Split(path, "/")
	for i, seg := range segments {
		segments[i] = echoSegmentToOpenAPI(seg)
	}
	return strings.Join(segments, "/")
}
