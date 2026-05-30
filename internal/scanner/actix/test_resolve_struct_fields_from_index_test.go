//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveStructFields_FromIndex 테스트
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveStructFields_FromIndex(t *testing.T) {

	src := []byte(`
struct Other { z: i64 }
struct User { id: i64, name: String }
`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: src, root: root}
	idx := buildStructIndex([]*fileInfo{fi})
	cache := map[string][]scanner.Field{}

	fields := resolveStructFields("User", idx, cache)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}

	if _, ok := cache["User"]; !ok {
		t.Error("expected User cached")
	}
}
