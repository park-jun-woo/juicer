//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_NoArgsCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandlePathParam_NoArgsCov(t *testing.T) {
	ep := &Endpoint{}
	handlePathParam(ep, &ast.CallExpr{})
}
