//ff:func feature=scan type=test control=sequence topic=actix
//ff:what isJSONFieldExpr — .json field_expression 판별 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestIsJSONFieldExpr_True(t *testing.T) {
	src := []byte(`fn f() { HttpResponse::Ok().json(x); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	// the field_expression "HttpResponse::Ok().json"
	fe := findCallByFuncSuffix(root, src, ".json")
	if fe == nil {
		t.Fatal("no .json call")
	}
	field := findChildByType(fe, "field_expression")
	if !isJSONFieldExpr(field, src) {
		t.Fatal("expected isJSONFieldExpr true")
	}
}

func TestIsJSONFieldExpr_OtherMethod(t *testing.T) {
	src := []byte(`fn f() { x.finish(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".finish")
	field := findChildByType(call, "field_expression")
	if isJSONFieldExpr(field, src) {
		t.Fatal("expected false for non-json field")
	}
}

func TestIsJSONFieldExpr_NotFieldExpr(t *testing.T) {
	// nil node and a non-field_expression node both yield false.
	if isJSONFieldExpr(nil, nil) {
		t.Fatal("expected false for nil node")
	}
	src := []byte(`fn f() { y; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var id *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if id != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "y" {
			id = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if isJSONFieldExpr(id, src) {
		t.Fatal("expected false for identifier node")
	}
}
