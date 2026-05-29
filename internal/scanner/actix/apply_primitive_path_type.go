//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 원시 타입 단일 path 파라미터의 타입을 설정한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyPrimitivePathType(ep *scanner.Endpoint, oaType string) {
	ensureRequest(ep)
	for i := range ep.Request.PathParams {
		ep.Request.PathParams[i].Type = oaType
	}
}
