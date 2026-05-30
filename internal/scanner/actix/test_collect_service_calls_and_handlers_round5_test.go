//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCollectServiceCallsAndHandlers_Round5 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectServiceCallsAndHandlers_Round5(t *testing.T) {
	fi := aFi(t, `fn config(cfg: &mut web::ServiceConfig) { cfg.service(web::scope("/api").service(h1).service(h2)); }`)
	var routes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectServiceCalls(n, fi.src, "/api", &routes)
		collectTopLevelServiceCall(n, fi, &routes)
	})

	scopeCall := aFirst(t, fi.root, "call_expression")
	_ = collectServiceHandlers(scopeCall, fi.src)

	_ = appendServiceCallHandlers(fi.root, fi.src, nil)
}
