//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what collectArrayElements 테스트
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

func TestCollectArrayElements_Empty(t *testing.T) {
	fi := mustParse(t, []byte("const x = [];\n"))
	arr := findAllByType(fi.Root, "array")[0]
	if got := collectArrayElements(arr); len(got) != 0 {
		t.Fatalf("expected 0 elements, got %d", len(got))
	}
}
