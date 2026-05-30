//ff:func feature=scan type=test control=sequence
//ff:what bindVarName — 바인딩 인자 변수명 추출 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestBindVarName(t *testing.T) {
	// &req -> "req"
	e1, err := parser.ParseExpr("&req")
	if err != nil {
		t.Fatal(err)
	}
	if got := bindVarName(e1); got != "req" {
		t.Errorf("&req -> %q, want req", got)
	}

	// plain identifier -> itself
	e2, _ := parser.ParseExpr("body")
	if got := bindVarName(e2); got != "body" {
		t.Errorf("body -> %q", got)
	}

	// non-AND unary (e.g. -x) falls through to exprString of whole expr
	e3, _ := parser.ParseExpr("-x")
	if got := bindVarName(e3); got == "" {
		t.Errorf("-x -> empty (unexpected)")
	}
}
