//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 경로 문자열에서 path 파라미터를 추출해 엔드포인트 요청에 설정한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyPathParams(ep *scanner.Endpoint, path string) {
	pathParams := extractPathParams(path)
	if len(pathParams) == 0 {
		return
	}
	ensureRequest(ep)
	ep.Request.PathParams = pathParams
}
