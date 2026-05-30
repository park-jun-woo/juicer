//ff:func feature=scan type=test control=sequence topic=actix
//ff:what indexFileStructs — 파일의 struct를 타입명 인덱스에 등록함을 검증
package actix

import "testing"

func TestIndexFileStructs(t *testing.T) {
	src := []byte(`
struct User { id: i64 }
struct Order { id: i64 }
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root}
	idx := make(structIndex)
	indexFileStructs(fi, idx)

	if len(idx) != 2 {
		t.Fatalf("expected 2 indexed structs, got %d", len(idx))
	}
	if e := idx["User"]; e == nil || e.structName != "User" || e.file != fi {
		t.Errorf("User entry incorrect: %+v", e)
	}
	if _, ok := idx["Order"]; !ok {
		t.Errorf("Order not indexed")
	}
}
