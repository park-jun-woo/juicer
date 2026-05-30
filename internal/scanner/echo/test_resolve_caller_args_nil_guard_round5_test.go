//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveCallerArgs_NilGuard_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestResolveCallerArgs_NilGuard_Round5(t *testing.T) {
	src := `package m
func Target() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)

	status, tn, fields, conf := resolveCallerArgs(fn, &ast.CallExpr{}, nil, nil)
	if status != "" || tn != "" || fields != nil || conf != "" {
		t.Fatalf("expected empty: %q %q %v %q", status, tn, fields, conf)
	}
}
