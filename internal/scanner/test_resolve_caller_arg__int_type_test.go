//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveCallerArg_IntType 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_IntType(t *testing.T) {
	ty := types.Typ[types.Int]
	r := resolveCallerArg(ty, &ast.BasicLit{Value: "200"}, nil)
	_ = r
}
