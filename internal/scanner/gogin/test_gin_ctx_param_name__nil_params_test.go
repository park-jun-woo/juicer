//ff:func feature=scan type=extract control=sequence
//ff:what TestGinCtxParamName_NilParams 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_NilParams(t *testing.T) {
	ft := &ast.FuncType{}
	got := ginCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
