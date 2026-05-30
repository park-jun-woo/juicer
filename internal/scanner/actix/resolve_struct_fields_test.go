//ff:func feature=scan type=test control=sequence topic=actix
//ff:what resolveStructFields — 캐시/인덱스/이름불일치 분기를 검증
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveStructFields_FromIndex(t *testing.T) {
	// Two structs in one file: the loop skips Other (continue) then matches User.
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
	// Now cached.
	if _, ok := cache["User"]; !ok {
		t.Error("expected User cached")
	}
}

func TestResolveStructFields_CacheHit(t *testing.T) {
	cache := map[string][]scanner.Field{
		"X": {{Name: "a", Type: "string"}},
	}
	fields := resolveStructFields("X", nil, cache)
	if len(fields) != 1 || fields[0].Name != "a" {
		t.Fatalf("expected cached fields, got %+v", fields)
	}
}

func TestResolveStructFields_NotInIndex(t *testing.T) {
	idx := structIndex{}
	if f := resolveStructFields("Missing", idx, map[string][]scanner.Field{}); f != nil {
		t.Fatalf("expected nil for unknown type, got %+v", f)
	}
}

func TestResolveStructFields_IndexedButNoMatch(t *testing.T) {
	// Index entry references a file whose structs do NOT include the type name,
	// so the loop never matches and the trailing `return nil` is reached.
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
