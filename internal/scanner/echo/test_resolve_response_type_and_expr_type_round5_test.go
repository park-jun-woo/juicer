//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveResponseType_And_ExprType_Round5 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_And_ExprType_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
type UserDto struct { Name string `+"`json:\"name\"`"+` }
var u UserDto
var _ = u
`)
	// find the use of u
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "u" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use of u")
	}
	name, fields := resolveExprType(useIdent, info)
	if name != "UserDto" || len(fields) == 0 {
		t.Fatalf("exprType: %q %+v", name, fields)
	}
	tn, flds, conf := resolveResponseType(useIdent, info)
	if tn != "UserDto" || len(flds) == 0 || conf != "full" {
		t.Fatalf("responseType: %q %+v %q", tn, flds, conf)
	}

	if n, _, _ := resolveResponseType(useIdent, nil); n != "" {
		t.Fatalf("nil info: %q", n)
	}
}
