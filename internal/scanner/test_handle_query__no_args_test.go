//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleQuery_NoArgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleQuery_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handleQuery(ep, call, "Query")
}
