//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestCollectArrayElements 테스트
package fastify

import "testing"

func TestCollectArrayElements(t *testing.T) {
	fi := mustParse(t, []byte(`const x = [1, "two", three];`+"\n"))
	arrs := findAllByType(fi.Root, "array")
	if len(arrs) == 0 {
		t.Fatal("no array node")
	}
	elems := collectArrayElements(arrs[0])
	if len(elems) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(elems))
	}
	for _, e := range elems {
		switch e.Type() {
		case "[", "]", ",":
			t.Errorf("separator leaked: %s", e.Type())
		}
	}
}
