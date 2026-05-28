//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 재스캔 결과의 엔드포인트를 (file, line) 기준으로 기존 endpoints에 반영한다
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyRescanResults(eps []scanner.Endpoint, ctx *groupArgCtx) {
	for _, ep := range eps {
		key := struct{ file string; line int }{ep.File, ep.Line}
		i, ok := ctx.epIndex[key]
		if !ok {
			continue
		}
		ctx.endpoints[i].Path = ep.Path
		if len(ep.Middleware) > 0 {
			ctx.endpoints[i].Middleware = ep.Middleware
		}
	}
}
