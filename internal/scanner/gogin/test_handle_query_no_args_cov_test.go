//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_NoArgsCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_NoArgsCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	handleQuery(ep, &ast.CallExpr{}, "Query")
}
