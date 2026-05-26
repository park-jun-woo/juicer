//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what valueAfterEquals 테스트
package fastapi

import "testing"

func TestValueAfterEquals(t *testing.T) {
	src := []byte("f(x=42)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kws := findAllByType(root, "keyword_argument")
	if len(kws) == 0 {
		t.Fatal("no keyword_argument")
	}
	got := valueAfterEquals(kws[0], src)
	if got != "42" {
		t.Fatalf("expected '42', got %q", got)
	}
}
