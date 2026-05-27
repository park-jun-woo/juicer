//ff:func feature=scan type=extract control=sequence
//ff:what TestHandlePathParam_NoArgs 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandlePathParam_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{}
	handlePathParam(ep, call)
}
