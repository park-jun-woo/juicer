//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryMatchKeyword_ValueAfterEquals 테스트
package fastapi

import "testing"

func TestTryMatchKeyword_ValueAfterEquals(t *testing.T) {

	src := []byte("f(prefix=PREFIX_CONST)\n")
	root, _ := parsePython(src)
	kws := findAllByType(root, "keyword_argument")
	if len(kws) == 0 {
		t.Fatal("no keyword_argument")
	}
	got := tryMatchKeyword(kws[0], "prefix", src)
	if got != "PREFIX_CONST" {
		t.Fatalf("valueAfterEquals: got %q, want PREFIX_CONST", got)
	}
}
