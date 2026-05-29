//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what struct 필드들을 path 파라미터 목록으로 펼쳐 설정한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyStructPathParams(ep *scanner.Endpoint, fields []scanner.Field) {
	ensureRequest(ep)
	var params []scanner.Param
	for _, f := range fields {
		params = append(params, scanner.Param{Name: f.JSON, Type: f.Type})
	}
	ep.Request.PathParams = params
}
