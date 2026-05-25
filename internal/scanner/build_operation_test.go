//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildOperation_Basic 테스트
package scanner

import "testing"

func TestBuildOperation_Basic(t *testing.T) {
	ep := Endpoint{Method: "GET", Path: "/api/health"}
	schemas := map[string]any{}
	op := buildOperation(ep, schemas)
	if op == nil {
		t.Fatal("expected non-nil")
	}
}
