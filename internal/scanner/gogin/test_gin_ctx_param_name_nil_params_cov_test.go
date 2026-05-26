//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamName_NilParamsCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinCtxParamName_NilParamsCov(t *testing.T) {
	ft := &ast.FuncType{}
	got := ginCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
