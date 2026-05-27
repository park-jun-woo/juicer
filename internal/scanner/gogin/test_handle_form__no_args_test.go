//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleForm_NoArgs 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleForm_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{}
	handleForm(ep, call)
}
