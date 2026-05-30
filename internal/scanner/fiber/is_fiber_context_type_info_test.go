//ff:func feature=scan type=test control=sequence
//ff:what isFiberContextTypeInfo — types.Type 기반 *fiber.Ctx 판정 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// typeOfVar type-checks src and returns the type of the named package-level var.
func typeOfVar(t *testing.T, src, varName string) types.Type {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Defs: map[*ast.Ident]types.Object{}}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	obj := pkg.Scope().Lookup(varName)
	if obj == nil {
		t.Fatalf("var %s not found", varName)
	}
	return obj.Type()
}

func TestIsFiberContextTypeInfo(t *testing.T) {
	src := `package m
type Ctx struct{}
type Other struct{}
var P = &Ctx{}      // pointer to named "Ctx" but pkg path is "m" (not fiber)
var O = &Other{}    // pointer to named "Other"
var I = 5           // not a pointer
var PB = new(int)   // pointer to non-named (basic)
`
	// *m.Ctx -> name is Ctx but pkg path "m" not fiber -> false
	if isFiberContextTypeInfo(typeOfVar(t, src, "P")) {
		t.Error("local *Ctx (non-fiber pkg) should be false")
	}
	// *Other -> name not Ctx -> false
	if isFiberContextTypeInfo(typeOfVar(t, src, "O")) {
		t.Error("*Other should be false")
	}
	// int -> not a pointer -> false
	if isFiberContextTypeInfo(typeOfVar(t, src, "I")) {
		t.Error("int should be false")
	}
	// *int -> pointer to basic (non-named) -> false
	if isFiberContextTypeInfo(typeOfVar(t, src, "PB")) {
		t.Error("*int should be false")
	}
}

func TestIsFiberContextTypeInfo_NamedNonCtx(t *testing.T) {
	// pointer to a named type from a real imported (non-fiber) package
	src := `package m
import "bytes"
var B = &bytes.Buffer{}
`
	if isFiberContextTypeInfo(typeOfVar(t, src, "B")) {
		t.Error("*bytes.Buffer should be false")
	}
}
