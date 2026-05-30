//ff:func feature=scan type=test control=selection
//ff:what inferValueType — 값 타입 추론 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func inferFor(t *testing.T, expr string) string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return inferValueType(e, nil)
}

func TestInferValueType_Literals(t *testing.T) {
	cases := map[string]string{
		`"x"`:      "string",
		"42":       "integer",
		"3.14":     "number",
		"true":     "boolean",
		"false":    "boolean",
		"nil":      "null",
		"[]int{1}": "array",
		"s[1:2]":   "array",
		"foo()":    "unknown", // not matched, nil info -> unknown
	}
	for in, want := range cases {
		if got := inferFor(t, in); got != want {
			t.Errorf("inferValueType(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestInferValueType_CompositeNonArray(t *testing.T) {
	// struct composite literal (not array) -> unknown
	if got := inferFor(t, "Book{}"); got != "unknown" {
		t.Errorf("struct composite -> %q, want unknown", got)
	}
}

func TestInferValueType_TypesFallback(t *testing.T) {
	// A type-checked expression resolves via info.Types fallback.
	src := `package m
var V = len("abc")
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	// find the len(...) call expr
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			call = c
			return false
		}
		return true
	})
	got := inferValueType(call, info)
	if got != "int" {
		t.Fatalf("types fallback: got %q, want int", got)
	}
}
