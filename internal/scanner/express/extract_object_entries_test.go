//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractObjectEntries: 유효 entry 추출 + nil entry 스킵 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstArray(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	arrs := findAllByType(fi.Root, "array")
	if len(arrs) == 0 {
		t.Fatal("no array")
	}
	return arrs[0]
}

func TestExtractObjectEntries_ValidAndSkip(t *testing.T) {
	// first object valid; second missing route -> nil -> skipped
	fi := mustParse(t, []byte(`const m = [ { path: '/a', route: userRoute }, { path: '/b' } ];`))
	entries := extractObjectEntries(firstArray(t, fi), fi.Src)
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d (%v)", len(entries), entries)
	}
	if entries[0].path != "/a" || entries[0].routeVar != "userRoute" {
		t.Fatalf("unexpected entry %+v", entries[0])
	}
}

func TestExtractObjectEntries_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`const m = [];`))
	entries := extractObjectEntries(firstArray(t, fi), fi.Src)
	if len(entries) != 0 {
		t.Fatalf("expected 0, got %d", len(entries))
	}
}
