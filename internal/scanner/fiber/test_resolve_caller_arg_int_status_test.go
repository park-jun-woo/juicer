//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_IntStatus 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArg_IntStatus(t *testing.T) {

	intType := types.Typ[types.Int]
	arg := &ast.BasicLit{Kind: token.INT, Value: "201"}
	res := resolveCallerArg(intType, arg, newEmptyInfo())
	if res.status != "201" {
		t.Fatalf("expected status 201, got %+v", res)
	}
}
