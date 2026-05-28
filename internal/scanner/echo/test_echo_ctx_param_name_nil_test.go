//ff:func feature=scan type=test control=sequence
//ff:what TestEchoCtxParamName_Nil 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoCtxParamName_Nil(t *testing.T) {
	ft := &ast.FuncType{Params: nil}
	got := echoCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
