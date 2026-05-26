//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_NoArgsCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleQuery_NoArgsCov(t *testing.T) {
	ep := &Endpoint{}
	handleQuery(ep, &ast.CallExpr{}, "Query")
}
