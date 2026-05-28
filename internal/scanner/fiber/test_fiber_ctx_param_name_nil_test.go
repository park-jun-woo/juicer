//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamName_NilParams 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamName_NilParams(t *testing.T) {
	ft := &ast.FuncType{}
	got := fiberCtxParamName(ft)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
