//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_NoArgsCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandlePathParam_NoArgsCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	handlePathParam(ep, &ast.CallExpr{})
}
