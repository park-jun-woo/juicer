//ff:func feature=scan type=extract control=sequence
//ff:what TestHandlePathParam_NoArgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandlePathParam_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handlePathParam(ep, call)
}
