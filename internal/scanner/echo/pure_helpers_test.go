//ff:func feature=scan type=test control=sequence topic=echo
//ff:what echo 순수/AST 헬퍼 함수 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func parseExpr(t *testing.T, code string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(code)
	if err != nil {
		t.Fatal(err)
	}
	return e
}

func TestEchoPathToOpenAPI(t *testing.T) {
	if got := echoPathToOpenAPI("/users/:id"); got != "/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := echoPathToOpenAPI("/files/*"); got != "/files/{wildcard}" {
		t.Fatalf("got %q", got)
	}
	if got := echoPathToOpenAPI("/files/*path"); got != "/files/{path}" {
		t.Fatalf("got %q", got)
	}
}

func TestIdentName(t *testing.T) {
	if identName(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if identName(parseExpr(t, "a.b")) != "" {
		t.Fatal("selector should be empty")
	}
}

func TestExprName(t *testing.T) {
	if exprName(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if exprName(parseExpr(t, "pkg.Fn")) != "pkg.Fn" {
		t.Fatal("selector")
	}
	if exprName(parseExpr(t, "fn()")) != "fn()" {
		t.Fatal("call")
	}
	if exprName(parseExpr(t, "func(){}")) != "(inline)" {
		t.Fatal("funclit")
	}
}

func TestIsEchoContextType(t *testing.T) {
	if !isEchoContextType(parseExpr(t, "echo.Context")) {
		t.Fatal("expected true for echo.Context")
	}
	if isEchoContextType(parseExpr(t, "gin.Context")) {
		t.Fatal("gin.Context should be false")
	}
	if isEchoContextType(parseExpr(t, "Foo")) {
		t.Fatal("plain ident should be false")
	}
}

func TestIsEchoInit(t *testing.T) {
	sel := parseExpr(t, "e.New").(*ast.SelectorExpr)
	if !isEchoInit(sel, "e") {
		t.Fatal("expected true for e.New")
	}
	if isEchoInit(sel, "other") {
		t.Fatal("wrong alias should be false")
	}
	sel2 := parseExpr(t, "e.Start").(*ast.SelectorExpr)
	if isEchoInit(sel2, "e") {
		t.Fatal("e.Start should be false")
	}
}

func TestResolveStatusCode_Literal(t *testing.T) {
	if got := resolveStatusCode(parseExpr(t, "200"), nil); got != "200" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveStatusCode_UnknownNilInfo(t *testing.T) {
	if got := resolveStatusCode(parseExpr(t, "someVar"), nil); got != "(unknown)" {
		t.Fatalf("got %q", got)
	}
}

func TestIsIntKind(t *testing.T) {
	if !isIntKind(types.Int) || !isIntKind(types.Uint64) {
		t.Fatal("int kinds")
	}
	if isIntKind(types.String) || isIntKind(types.Float64) {
		t.Fatal("non-int kinds")
	}
}

func TestAstFieldJSONName(t *testing.T) {
	src := `package m
type S struct {
	Name string ` + "`json:\"name\"`" + `
	Skip string ` + "`json:\"-\"`" + `
	Plain string
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var fields []*ast.Field
	ast.Inspect(file, func(n ast.Node) bool {
		if st, ok := n.(*ast.StructType); ok {
			fields = st.Fields.List
		}
		return true
	})
	if len(fields) < 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}
	if got := astFieldJSONName(fields[0]); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := astFieldJSONName(fields[1]); got != "" {
		t.Fatalf("dash should be empty: %q", got)
	}
	if got := astFieldJSONName(fields[2]); got != "" {
		t.Fatalf("no tag should be empty: %q", got)
	}
}
