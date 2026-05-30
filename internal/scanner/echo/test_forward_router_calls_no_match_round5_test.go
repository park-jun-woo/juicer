//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestForwardRouterCalls_NoMatch_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestForwardRouterCalls_NoMatch_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	src := `package m
func f() { other(); }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)

	forwardRouterCalls(fn.Body.List, "router", "/p", &routerInfo{}, nil, ctx, 0)
}
