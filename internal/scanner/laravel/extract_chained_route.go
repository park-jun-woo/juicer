//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::middleware(..)->get('/path', handler) 체인 단일 라우트를 routeInfo로 추출한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// extractChainedRoute handles a single chained route such as
// Route::middleware('auth')->prefix('v1')->get('/user', handler) where the
// terminal member call is an HTTP verb rather than group().
func extractChainedRoute(mc *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) *routeInfo {
	method := lastMemberCallName(mc, fi.src)
	upperMethod, ok := httpMethods[strings.ToLower(method)]
	if !ok {
		return nil
	}

	prefix := outerPrefix
	mw := copyMiddleware(outerMiddleware)
	if inner := findChildByType(mc, "scoped_call_expression"); inner != nil {
		walkChain(inner, fi, &prefix, &mw)
	} else if innerMC := findChildByType(mc, "member_call_expression"); innerMC != nil {
		walkChain(innerMC, fi, &prefix, &mw)
	} else {
		return nil
	}

	return chainedRouteInfo(mc, fi, upperMethod, prefix, mw)
}
