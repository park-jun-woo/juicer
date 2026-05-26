//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_NoArgsCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleForm_NoArgsCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	handleForm(ep, &ast.CallExpr{})
}
