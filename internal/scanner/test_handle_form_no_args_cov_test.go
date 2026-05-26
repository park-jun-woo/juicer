//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_NoArgsCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleForm_NoArgsCov(t *testing.T) {
	ep := &Endpoint{}
	handleForm(ep, &ast.CallExpr{})
}
