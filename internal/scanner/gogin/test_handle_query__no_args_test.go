//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleQuery_NoArgs 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{}
	handleQuery(ep, call, "Query")
}
