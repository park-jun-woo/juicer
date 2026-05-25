package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestScanBody_NilBodyCase(t *testing.T) {
	ep := &Endpoint{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	scanBody(ep, nil, "c", nil, idx, "handler")
}

func TestScanBody_EmptyBodyCase(t *testing.T) {
	ep := &Endpoint{}
	body := &ast.BlockStmt{}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	scanBody(ep, body, "c", nil, idx, "handler")
}
