//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractKeywordArg 테스트
package fastapi

import "testing"

func TestExtractKeywordArg(t *testing.T) {
	src := []byte(`f(prefix="/api", status_code=201)` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	argsList := findAllByType(root, "argument_list")
	if len(argsList) == 0 {
		t.Fatal("no argument_list")
	}
	got := extractKeywordArg(argsList[0], "prefix", src)
	if got != "/api" {
		t.Fatalf("prefix: expected '/api', got %q", got)
	}
	got2 := extractKeywordArg(argsList[0], "status_code", src)
	if got2 != "201" {
		t.Fatalf("status_code: expected '201', got %q", got2)
	}
	got3 := extractKeywordArg(argsList[0], "nonexistent", src)
	if got3 != "" {
		t.Fatalf("nonexistent: expected '', got %q", got3)
	}
}
