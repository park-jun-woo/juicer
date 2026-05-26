//ff:func feature=scan type=test control=sequence
//ff:what TestGinCtxParamNameInfo_NilParams 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestGinCtxParamNameInfo_NilParams(t *testing.T) {
	ft := &ast.FuncType{Params: nil}
	got := ginCtxParamNameInfo(ft, &types.Info{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
