//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestTryGroupArgCall_And_WalkForGroupArgs_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestTryGroupArgCall_And_WalkForGroupArgs_Round5(t *testing.T) {
	ctx := emptyGroupCtx()

	call := callExprFrom(t, `setup(42)`)
	tryGroupArgCall(call, ctx)

	src := `package m
func f() {
	x := 1
	g(x)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	walkForGroupArgs(fn.Body.List, ctx)
}
