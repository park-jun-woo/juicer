//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectArgNodes 테스트
package hono

import "testing"

func TestCollectArgNodes(t *testing.T) {
	fi := mustParse(t, []byte(`f("a", b, { c: 1 });`+"\n"))
	args := findAllByType(fi.Root, "arguments")
	if len(args) == 0 {
		t.Fatal("no arguments")
	}
	nodes := collectArgNodes(args[0])
	if len(nodes) != 3 {
		t.Fatalf("expected 3 arg nodes, got %d", len(nodes))
	}
	for _, n := range nodes {
		switch n.Type() {
		case "(", ")", ",":
			t.Errorf("separator leaked: %s", n.Type())
		}
	}
}
