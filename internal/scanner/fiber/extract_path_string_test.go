//ff:func feature=scan type=test control=selection
//ff:what extractPathString — 경로 문자열 추출 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func extractPathFor(t *testing.T, expr string) (string, bool) {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return extractPathString(e)
}

func TestExtractPathString_Literal(t *testing.T) {
	got, ok := extractPathFor(t, `"/users"`)
	if !ok || got != "/users" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

func TestExtractPathString_Concat(t *testing.T) {
	got, ok := extractPathFor(t, `"/api" + "/v1"`)
	if !ok || got != "/api/v1" {
		t.Fatalf("got %q ok=%v", got, ok)
	}
}

func TestExtractPathString_NonString(t *testing.T) {
	_, ok := extractPathFor(t, "42")
	if ok {
		t.Fatal("int literal should be false")
	}
}

func TestExtractPathString_Other(t *testing.T) {
	// an identifier matches no case -> false
	_, ok := extractPathFor(t, "pathVar")
	if ok {
		t.Fatal("ident should be false")
	}
}
