//ff:func feature=scan type=test control=selection
//ff:what collectStringParts — 문자열 연결 파트 수집 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func collectFor(t *testing.T, expr string) []string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	var parts []string
	collectStringParts(e, &parts)
	return parts
}

func TestCollectStringParts_SingleString(t *testing.T) {
	if got := collectFor(t, `"hello"`); len(got) != 1 || got[0] != "hello" {
		t.Fatalf("got %v", got)
	}
}

func TestCollectStringParts_Concat(t *testing.T) {
	// "/api" + "/v1" + baseURL (ident skipped) + opts.X (selector skipped)
	got := collectFor(t, `"/api" + "/v1" + baseURL + opts.X`)
	if len(got) != 2 || got[0] != "/api" || got[1] != "/v1" {
		t.Fatalf("got %v, want [/api /v1]", got)
	}
}

func TestCollectStringParts_NonStringLit(t *testing.T) {
	// integer literal -> nothing collected
	if got := collectFor(t, "42"); len(got) != 0 {
		t.Fatalf("expected empty for int lit, got %v", got)
	}
}

func TestCollectStringParts_NonAddBinary(t *testing.T) {
	// subtraction -> not traversed
	if got := collectFor(t, "a - b"); len(got) != 0 {
		t.Fatalf("expected empty for non-ADD binary, got %v", got)
	}
}

func TestCollectStringParts_BareIdent(t *testing.T) {
	if got := collectFor(t, "baseURL"); len(got) != 0 {
		t.Fatalf("expected empty for ident, got %v", got)
	}
}
