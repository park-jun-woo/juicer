//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what nodeText 테스트
package fastify

import "testing"

func TestNodeText(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	id := findAllByType(fi.Root, "identifier")
	if len(id) == 0 {
		t.Fatal("no identifier")
	}
	if got := nodeText(id[0], fi.Src); got != "x" {
		t.Fatalf("nodeText = %q, want x", got)
	}
}
