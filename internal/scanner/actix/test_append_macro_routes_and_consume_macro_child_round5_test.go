//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestAppendMacroRoutes_And_ConsumeMacroChild_Round5 테스트
package actix

import "testing"

func TestAppendMacroRoutes_And_ConsumeMacroChild_Round5(t *testing.T) {
	fi := aFi(t, macroSrc)
	fn := aFirst(t, fi.root, "function_item")
	var routes, pending []macroRoute

	pending = append(pending, macroRoute{method: "GET", path: "/health"})
	routes = appendMacroRoutes(routes, pending, fn, fi, "health")
	if len(routes) == 0 {
		t.Fatalf("expected macro routes attached")
	}

	// consumeMacroChild over the top-level children
	var r2, p2 []macroRoute
	for i := 0; i < int(fi.root.ChildCount()); i++ {
		r2, p2 = consumeMacroChild(fi.root.Child(i), fi, r2, p2)
	}
}
