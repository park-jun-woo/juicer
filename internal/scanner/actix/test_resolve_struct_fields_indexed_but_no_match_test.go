//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveStructFields_IndexedButNoMatch 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveStructFields_IndexedButNoMatch(t *testing.T) {

	src := []byte(`struct Other { z: i64 }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root}
	idx := structIndex{
		"Ghost": &structEntry{file: fi, structName: "Ghost"},
	}
	if f := resolveStructFields("Ghost", idx, map[string][]scanner.Field{}); f != nil {
		t.Fatalf("expected nil when no struct_item matches, got %+v", f)
	}
}
