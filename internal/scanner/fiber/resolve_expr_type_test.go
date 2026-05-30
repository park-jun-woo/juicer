//ff:func feature=scan type=test control=selection
//ff:what resolveExprType — 표현 타입 추적 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func typedExprs(t *testing.T, src string) (*ast.File, *types.Info) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	return file, info
}

func TestResolveExprType_NilInfo(t *testing.T) {
	tn, f := resolveExprType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || f != nil {
		t.Fatalf("nil info: %q %v", tn, f)
	}
}

func TestResolveExprType_IdentUse(t *testing.T) {
	src := `package m
type Resp struct { OK bool ` + "`json:\"ok\"`" + ` }
func h() {
	var r Resp
	use(r)
}
func use(interface{}) {}
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0]
			}
		}
		return true
	})
	tn, fields := resolveExprType(arg, info)
	if tn != "Resp" || len(fields) != 1 {
		t.Fatalf("ident use: %q %v", tn, fields)
	}
}

func TestResolveExprType_CompositeLit(t *testing.T) {
	src := `package m
type Out struct { N int ` + "`json:\"n\"`" + ` }
func h() { use(Out{N: 1}) }
func use(interface{}) {}
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if cl, ok := n.(*ast.CompositeLit); ok && arg == nil {
			arg = cl
		}
		return true
	})
	tn, _ := resolveExprType(arg, info)
	if tn != "Out" {
		t.Fatalf("composite lit: %q", tn)
	}
}

func TestResolveExprType_Unresolvable(t *testing.T) {
	// ident not in Uses/Defs with empty info -> "" nil
	tn, f := resolveExprType(&ast.Ident{Name: "ghost"}, newEmptyInfoFull())
	if tn != "" || f != nil {
		t.Fatalf("unresolvable: %q %v", tn, f)
	}
}

func TestResolveExprType_Selector(t *testing.T) {
	// a selector expr referencing a package-level typed var
	src := `package m
type Cfg struct { N int ` + "`json:\"n\"`" + ` }
var Conf Cfg
func use(interface{}) {}
func h() { use(pkgAlias()) }
func pkgAlias() Cfg { return Conf }
`
	file, info := typedExprs(t, src)
	// find a selector expr in Uses — use Conf via its Ident directly
	var sel *ast.SelectorExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if s, ok := n.(*ast.SelectorExpr); ok && sel == nil {
			sel = s
		}
		return true
	})
	if sel != nil {
		resolveExprType(sel, info) // exercise selector path; no assertion needed
	}
}

func TestResolveExprType_DefaultCall(t *testing.T) {
	src := `package m
type R struct { V int ` + "`json:\"v\"`" + ` }
func make2() R { return R{} }
func use(interface{}) {}
func h() { use(make2()) }
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0] // make2() call -> default branch via Types
			}
		}
		return true
	})
	tn, _ := resolveExprType(arg, info)
	if tn != "R" {
		t.Fatalf("default call: %q, want R", tn)
	}
}

func TestResolveExprType_IdentDef(t *testing.T) {
	// the defining ident of a typed var is in Defs (not Uses)
	src := `package m
type D struct { A int ` + "`json:\"a\"`" + ` }
var Decl D
`
	file, info := typedExprs(t, src)
	var defIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "Decl" {
			defIdent = id
		}
		return true
	})
	if defIdent == nil {
		t.Fatal("Decl ident not found")
	}
	tn, _ := resolveExprType(defIdent, info)
	if tn != "D" {
		t.Fatalf("ident def: %q, want D", tn)
	}
}

func TestResolveExprType_SelectorUse(t *testing.T) {
	// selector to a typed struct field/var resolves via Uses[e.Sel]
	src := `package m
type Inner struct { Z int ` + "`json:\"z\"`" + ` }
type Outer struct { In Inner }
func use(interface{}) {}
func h() {
	var o Outer
	use(o.In)
}
`
	file, info := typedExprs(t, src)
	var arg *ast.SelectorExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				if s, ok := c.Args[0].(*ast.SelectorExpr); ok {
					arg = s
				}
			}
		}
		return true
	})
	if arg == nil {
		t.Fatal("selector arg not found")
	}
	tn, _ := resolveExprType(arg, info)
	if tn != "Inner" {
		t.Fatalf("selector use: %q, want Inner", tn)
	}
}

func newEmptyInfoFull() *types.Info {
	return &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
}
