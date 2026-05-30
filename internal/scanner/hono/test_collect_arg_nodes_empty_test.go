//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestCollectArgNodes_Empty 테스트
package hono

import "testing"

func TestCollectArgNodes_Empty(t *testing.T) {
	fi := mustParse(t, []byte("f();\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	if got := collectArgNodes(args); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}
