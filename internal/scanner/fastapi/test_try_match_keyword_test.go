//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryMatchKeyword 테스트
package fastapi

import "testing"

func TestTryMatchKeyword(t *testing.T) {
	src := []byte(`f(prefix="/api", status_code=201)` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kws := findAllByType(root, "keyword_argument")
	if len(kws) < 2 {
		t.Fatalf("expected >= 2 keyword_arguments, got %d", len(kws))
	}

	got := tryMatchKeyword(kws[0], "prefix", src)
	if got != "/api" {
		t.Fatalf("prefix: got %q", got)
	}

	got2 := tryMatchKeyword(kws[0], "wrong_name", src)
	if got2 != "" {
		t.Fatalf("wrong_name: expected empty, got %q", got2)
	}

	got3 := tryMatchKeyword(kws[1], "status_code", src)
	if got3 != "201" {
		t.Fatalf("status_code: got %q", got3)
	}
}
