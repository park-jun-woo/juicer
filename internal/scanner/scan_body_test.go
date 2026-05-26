//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_NilBodyCase 테스트
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
