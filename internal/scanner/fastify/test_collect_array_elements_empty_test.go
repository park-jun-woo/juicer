//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectArrayElements_Empty 테스트
package fastify

import "testing"

func TestCollectArrayElements_Empty(t *testing.T) {
	fi := mustParse(t, []byte("const x = [];\n"))
	arr := findAllByType(fi.Root, "array")[0]
	if got := collectArrayElements(arr); len(got) != 0 {
		t.Fatalf("expected 0 elements, got %d", len(got))
	}
}
