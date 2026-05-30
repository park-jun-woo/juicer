//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_OtherType 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArg_OtherType(t *testing.T) {

	res := resolveCallerArg(types.Typ[types.String], &ast.BasicLit{Kind: token.STRING, Value: `"x"`}, newEmptyInfo())
	if res.status != "" || res.typeName != "" {
		t.Fatalf("expected empty result for string param, got %+v", res)
	}
}
