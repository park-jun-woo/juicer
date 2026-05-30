//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractPairIdentValue: 키매칭+ident값 / 키불일치 / ident아님 분기
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstObject(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object")
	}
	return objs[0]
}

func TestExtractPairIdentValue_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { route: userRoute, other: x };`))
	if got := extractPairIdentValue(firstObject(t, fi), fi.Src, "route"); got != "userRoute" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPairIdentValue_KeyMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { foo: bar };`))
	if got := extractPairIdentValue(firstObject(t, fi), fi.Src, "route"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractPairIdentValue_ValueNotIdent(t *testing.T) {
	// matching key but value is a string, not an identifier
	fi := mustParse(t, []byte(`const o = { route: '/literal' };`))
	if got := extractPairIdentValue(firstObject(t, fi), fi.Src, "route"); got != "" {
		t.Fatalf("got %q", got)
	}
}
