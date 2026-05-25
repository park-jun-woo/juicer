//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleForm_NoArgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleForm_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handleForm(ep, call)
}
