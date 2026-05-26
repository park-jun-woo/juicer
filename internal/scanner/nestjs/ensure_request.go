//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 엔드포인트에 Request가 없으면 빈 Request를 생성한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// ensureRequest creates an empty Request on the endpoint if nil.
func ensureRequest(ep *scanner.Endpoint) {
	if ep.Request == nil {
		ep.Request = &scanner.Request{}
	}
}
