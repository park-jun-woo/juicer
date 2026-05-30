//ff:func feature=scan type=test control=sequence
//ff:what TestFiberCtxParamNameInfo_NilParams 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberCtxParamNameInfo_NilParams(t *testing.T) {
	if got := fiberCtxParamNameInfo(&ast.FuncType{}, newEmptyInfo()); got != "" {
		t.Fatalf("nil params: got %q", got)
	}
}
