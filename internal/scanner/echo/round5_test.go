//ff:func feature=scan type=test control=sequence topic=echo
//ff:what round5 미커버 순수/types 헬퍼 직접 호출 테스트 (echo)
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// checkSrc type-checks a small in-memory package and returns the file + info.
func checkSrc(t *testing.T, src string) (*ast.File, *types.Info) {
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

func firstExprOfType[T ast.Node](t *testing.T, file *ast.File, pred func(T) bool) T {
	t.Helper()
	var found T
	var ok bool
	ast.Inspect(file, func(n ast.Node) bool {
		if ok {
			return false
		}
		if tn, is := n.(T); is && pred(tn) {
			found = tn
			ok = true
			return false
		}
		return true
	})
	if !ok {
		t.Fatal("node not found")
	}
	return found
}

func TestUnquote_Round5(t *testing.T) {
	if unquote(`"hello"`) != "hello" {
		t.Fatal("double-quote")
	}
	if unquote("`raw`") != "raw" {
		t.Fatal("raw")
	}
	if unquote("plain") != "plain" {
		t.Fatal("plain")
	}
}

func TestSlicePrefix_Round5(t *testing.T) {
	if slicePrefix(true) != "[]" {
		t.Fatal("slice")
	}
	if slicePrefix(false) != "" {
		t.Fatal("non-slice")
	}
}

func TestStringLitValue_Round5(t *testing.T) {
	if got := stringLitValue(parseExpr(t, `"abc"`)); got != "abc" {
		t.Fatalf("got %q", got)
	}
	if got := stringLitValue(parseExpr(t, `42`)); got != "" {
		t.Fatalf("non-string: %q", got)
	}
	if got := stringLitValue(parseExpr(t, `foo`)); got != "" {
		t.Fatalf("ident: %q", got)
	}
}

func TestFormatType_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var A int
var B *int
var C []string
var D map[string]int
var E [3]byte
`)
	want := map[string]string{"A": "int", "B": "*int", "C": "[]string", "D": "map[string]int", "E": "[]byte"}
	for name, exp := range want {
		var typ types.Type
		for id, obj := range info.Defs {
			if obj != nil && id.Name == name {
				typ = obj.Type()
			}
		}
		if typ == nil {
			t.Fatalf("no type for %s", name)
		}
		if got := formatType(typ); got != exp {
			t.Errorf("formatType(%s)=%q want %q", name, got, exp)
		}
	}
}

func TestUnwrapPointer_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var P *int
var I int
`)
	var pt, it types.Type
	for id, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if id.Name == "P" {
			pt = obj.Type()
		}
		if id.Name == "I" {
			it = obj.Type()
		}
	}
	if _, ok := unwrapPointer(pt).(*types.Basic); !ok {
		t.Fatalf("pointer not unwrapped: %v", unwrapPointer(pt))
	}
	if unwrapPointer(it) != it {
		t.Fatal("non-pointer should be returned unchanged")
	}
}

func TestInferValueType_Round5(t *testing.T) {
	cases := map[string]string{
		`"s"`:   "string",
		`42`:    "integer",
		`3.14`:  "number",
		`true`:  "boolean",
		`false`: "boolean",
		`nil`:   "null",
	}
	for src, want := range cases {
		if got := inferValueType(parseExpr(t, src), nil); got != want {
			t.Errorf("inferValueType(%s)=%q want %q", src, got, want)
		}
	}
}

func TestCollectStringParts_Round5(t *testing.T) {
	var parts []string
	collectStringParts(nil, parseExpr(t, `"a" + "b" + "c"`), &parts)
	if len(parts) != 3 || parts[0] != "a" || parts[2] != "c" {
		t.Fatalf("parts: %v", parts)
	}
}

func TestExtractPathString_Round5(t *testing.T) {
	got, ok := extractPathString(nil, parseExpr(t, `"/users"`))
	if !ok || got != "/users" {
		t.Fatalf("literal: %q %v", got, ok)
	}
	// concatenation of two literals
	got2, ok2 := extractPathString(nil, parseExpr(t, `"/a" + "/b"`))
	if !ok2 || got2 != "/a/b" {
		t.Fatalf("concat: %q %v", got2, ok2)
	}
	// non-resolvable
	if _, ok := extractPathString(nil, parseExpr(t, `someVar`)); ok {
		t.Fatal("ident should not resolve without info")
	}
}

func TestExtractBinaryPath_Round5(t *testing.T) {
	// "/a" + "/b" binary expr
	be := parseExpr(t, `"/a" + "/b"`).(*ast.BinaryExpr)
	got, ok := extractBinaryPath(nil, be)
	if !ok || got != "/a/b" {
		t.Fatalf("got %q %v", got, ok)
	}
}

func TestWellKnownType_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
import "time"
var T time.Time
`)
	var named *types.Named
	for id, obj := range info.Defs {
		if obj != nil && id.Name == "T" {
			if n, ok := obj.Type().(*types.Named); ok {
				named = n
			}
		}
	}
	if named == nil {
		t.Fatal("no named type")
	}
	name, ok := wellKnownType(named)
	if !ok || name != "time.Time" {
		t.Fatalf("got %q %v", name, ok)
	}
}
