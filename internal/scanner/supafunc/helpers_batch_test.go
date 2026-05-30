//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what supafunc AST/순수 헬퍼 함수 테스트
package supafunc

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestUnquoteTS(t *testing.T) {
	if unquoteTS(`"x"`) != "x" || unquoteTS("'y'") != "y" || unquoteTS("`z`") != "z" || unquoteTS("a") != "a" {
		t.Fatal("unquote")
	}
}

func TestNodeTextAndFindChild(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	if findChildByType(calls[0], "arguments") == nil {
		t.Fatal("arguments")
	}
	if findChildByType(calls[0], "object") != nil {
		t.Fatal("nil expected")
	}
	id := findChildByType(calls[0], "identifier")
	if id != nil && nodeText(id, fi.Src) != "foo" {
		t.Fatalf("name %q", nodeText(id, fi.Src))
	}
}

func TestChildrenOfType(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1, b: 2 };`))
	objs := findAllByType(fi.Root, "object")
	if len(childrenOfType(objs[0], "pair")) != 2 {
		t.Fatal("children")
	}
}

func TestFindAllByType(t *testing.T) {
	fi := mustParse(t, []byte(`a(); b(); c();`))
	if len(findAllByType(fi.Root, "call_expression")) != 3 {
		t.Fatal("findAll")
	}
}

func TestWalkNodes(t *testing.T) {
	fi := mustParse(t, []byte(`a(b());`))
	count := 0
	walkNodes(fi.Root, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("got %d", count)
	}
}

func TestOperatorOfBinaryExpr(t *testing.T) {
	fi := mustParse(t, []byte(`const x = a === b;`))
	bins := findAllByType(fi.Root, "binary_expression")
	if len(bins) == 0 {
		t.Fatal("no binary expr")
	}
	if got := operatorOfBinaryExpr(bins[0], fi.Src); got != "===" {
		t.Fatalf("got %q", got)
	}
}

func TestHasChildBinaryExpr(t *testing.T) {
	fi := mustParse(t, []byte(`if (a === b) {}`))
	parens := findAllByType(fi.Root, "parenthesized_expression")
	if len(parens) == 0 {
		t.Skip("no parenthesized expr")
	}
	if !hasChildBinaryExpr(parens[0]) {
		t.Fatal("expected binary child")
	}
	fi2 := mustParse(t, []byte(`if (a) {}`))
	parens2 := findAllByType(fi2.Root, "parenthesized_expression")
	if len(parens2) > 0 && hasChildBinaryExpr(parens2[0]) {
		t.Fatal("no binary child expected")
	}
}

func TestExtractStatusFromObject(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { status: 201, headers: {} };`))
	objs := findAllByType(fi.Root, "object")
	if got := extractStatusFromObject(objs[0], fi.Src); got != "201" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromObject_None(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { headers: {} };`))
	objs := findAllByType(fi.Root, "object")
	if got := extractStatusFromObject(objs[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "index.ts", `serve(async (req) => new Response("ok"));`)
	fi, err := parseFile(dir + "/index.ts")
	if err != nil || fi == nil || fi.Root == nil {
		t.Fatalf("err: %v", err)
	}
}

func TestParseFile_Missing(t *testing.T) {
	if _, err := parseFile("/no/such.ts"); err == nil {
		t.Fatal("expected error")
	}
}
