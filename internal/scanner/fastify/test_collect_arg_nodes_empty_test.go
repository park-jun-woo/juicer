//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectArgNodes_Empty 테스트
package fastify

import "testing"

func TestCollectArgNodes_Empty(t *testing.T) {
	fi := mustParse(t, []byte("f();\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	if got := collectArgNodes(args); len(got) != 0 {
		t.Fatalf("expected 0 nodes, got %d", len(got))
	}
}
