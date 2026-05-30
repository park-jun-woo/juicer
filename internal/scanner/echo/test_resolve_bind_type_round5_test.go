//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveBindType_Round5 테스트
package echo

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestResolveBindType_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
type CreateReq struct { Name string `+"`json:\"name\"`"+` }
var dto CreateReq
var _ = dto
`)
	// build a Bind call: c.Bind(&dto) referencing the used dto ident
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "dto" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use of dto")
	}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.UnaryExpr{Op: token.AND, X: useIdent}}}
	name, fields := resolveBindType(call, info)
	if name != "CreateReq" || len(fields) == 0 {
		t.Fatalf("bind type: %q %+v", name, fields)
	}

	if n, _ := resolveBindType(&ast.CallExpr{}, info); n != "" {
		t.Fatalf("no-args: %q", n)
	}

	if n, _ := resolveBindType(call, nil); n != "" {
		t.Fatalf("nil info: %q", n)
	}
}
