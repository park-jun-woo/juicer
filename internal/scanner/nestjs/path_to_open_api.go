//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what NestJS 경로 패턴(:param)을 OpenAPI 경로 패턴({param})으로 변환한다
package nestjs

import "strings"

// pathToOpenAPI converts NestJS route parameters (:id) to OpenAPI format ({id}).
func pathToOpenAPI(path string) string {
	segments := strings.Split(path, "/")
	for i, seg := range segments {
		if strings.HasPrefix(seg, ":") {
			segments[i] = "{" + seg[1:] + "}"
		}
	}
	return strings.Join(segments, "/")
}
