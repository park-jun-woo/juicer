//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestPropagateParent_NoPrefix 테스트
package express

import "testing"

func TestPropagateParent_NoPrefix(t *testing.T) {
	g := newTestGraph()
	parent := routerKey{file: "a", varName: "app"}
	if propagateParent(g, parent, map[routerKey][]string{}) {
		t.Fatal("expected false when parent has no prefix")
	}
}
