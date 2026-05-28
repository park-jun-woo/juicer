//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 엔드포인트에 Request가 없으면 생성한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func ensureRequest(ep *scanner.Endpoint) {
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
}
