//ff:func feature=scan type=test control=sequence topic=actix
//ff:what buildStructIndex — 여러 파일의 struct를 타입명 인덱스로 구축함을 검증
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

func TestBuildStructIndex_Empty(t *testing.T) {
	idx := buildStructIndex(nil)
	if len(idx) != 0 {
		t.Fatalf("expected empty index, got %d", len(idx))
	}
}
