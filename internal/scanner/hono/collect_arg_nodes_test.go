//ff:func feature=scan type=test control=sequence topic=hono
//ff:what collectArgNodes 테스트
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

func TestCollectArgNodes_Empty(t *testing.T) {
	fi := mustParse(t, []byte("f();\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	if got := collectArgNodes(args); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
