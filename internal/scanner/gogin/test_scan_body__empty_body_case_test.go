//ff:func feature=scan type=extract control=sequence
//ff:what TestScanBody_EmptyBodyCase 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestScanBody_EmptyBodyCase(t *testing.T) {
	ep := &scanner.Endpoint{}
	body := &ast.BlockStmt{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	scanBody(ep, body, "c", nil, idx, "handler")
}
