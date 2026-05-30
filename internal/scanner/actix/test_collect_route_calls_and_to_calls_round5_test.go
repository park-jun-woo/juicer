//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCollectRouteCalls_And_ToCalls_Round5 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectRouteCalls_And_ToCalls_Round5(t *testing.T) {
	fi := aFi(t, builderSrc)
	var routes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectRouteCalls(n, fi.src, "/users", &routes)
	})
	if len(routes) == 0 {
		t.Fatalf("collectRouteCalls found nothing")
	}

	var toRoutes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectToCalls(n, fi.src, "/users", &toRoutes)
	})
}
