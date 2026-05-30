//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestBuildStructIndex 테스트
package actix

import "testing"

func TestBuildStructIndex(t *testing.T) {
	srcA := []byte(`struct User { id: i64 }`)
	rootA, err := parseRust(srcA)
	if err != nil {
		t.Fatal(err)
	}
	srcB := []byte(`struct Order { id: i64 }`)
	rootB, err := parseRust(srcB)
	if err != nil {
		t.Fatal(err)
	}
	files := []*fileInfo{
		{src: srcA, root: rootA},
		{src: srcB, root: rootB},
	}
	idx := buildStructIndex(files)
	if len(idx) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(idx))
	}
	if _, ok := idx["User"]; !ok {
		t.Error("User missing from index")
	}
	if _, ok := idx["Order"]; !ok {
		t.Error("Order missing from index")
	}
}
