//ff:func feature=scan type=test control=sequence topic=hono
//ff:what AST 헬퍼(findChildByType, findAllByType, childrenOfType, findObjectValueByKey, walkNodes, nodeText) 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestFindChildByType_Found(t *testing.T) {
	fi := mustParse(t, []byte(`f();`+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if findChildByType(call, "arguments") == nil {
		t.Fatal("expected arguments child")
	}
}

func TestFindChildByType_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`f();`+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	if findChildByType(call, "object") != nil {
		t.Fatal("expected nil for missing type")
	}
}

func TestFindAllByType(t *testing.T) {
	fi := mustParse(t, []byte(`a(); b(); c();`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) != 3 {
		t.Fatalf("expected 3 calls, got %d", len(calls))
	}
}

func TestFindAllByType_None(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 1;`+"\n"))
	if got := findAllByType(fi.Root, "call_expression"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}

func TestChildrenOfType(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1, b: 2, c: 3 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	pairs := childrenOfType(obj, "pair")
	if len(pairs) != 3 {
		t.Fatalf("expected 3 pairs, got %d", len(pairs))
	}
}

func TestChildrenOfType_None(t *testing.T) {
	fi := mustParse(t, []byte(`const o = {};`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if got := childrenOfType(obj, "pair"); len(got) != 0 {
		t.Fatalf("expected 0, got %d", len(got))
	}
}

func TestFindObjectValueByKey_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { method: "get", path: "/x" };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	v := findObjectValueByKey(obj, "path", fi.Src)
	if v == nil || nodeText(v, fi.Src) != `"/x"` {
		t.Fatalf("got %v", v)
	}
}

func TestFindObjectValueByKey_NotFound(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { a: 1 };`+"\n"))
	obj := findAllByType(fi.Root, "object")[0]
	if findObjectValueByKey(obj, "missing", fi.Src) != nil {
		t.Fatal("expected nil")
	}
}

func TestFindObjectValueByKey_NilObj(t *testing.T) {
	if findObjectValueByKey(nil, "x", []byte("")) != nil {
		t.Fatal("expected nil for nil obj")
	}
}

func TestWalkNodes(t *testing.T) {
	fi := mustParse(t, []byte(`a(b());`+"\n"))
	count := 0
	walkNodes(fi.Root, func(n *sitter.Node) {
		if n.Type() == "call_expression" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("expected 2 call_expression visits, got %d", count)
	}
}

func TestNodeText(t *testing.T) {
	fi := mustParse(t, []byte(`hello;`+"\n"))
	id := findAllByType(fi.Root, "identifier")[0]
	if got := nodeText(id, fi.Src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}
