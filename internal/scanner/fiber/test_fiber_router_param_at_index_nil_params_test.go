//ff:func feature=scan type=test control=sequence
//ff:what TestFiberRouterParamAtIndex_NilParams 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberRouterParamAtIndex_NilParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	if got := fiberRouterParamAtIndex(fn, newEmptyInfo(), 0); got != "" {
		t.Fatalf("nil params: got %q", got)
	}
}
