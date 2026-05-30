//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryExtractDefault 분기 전수 테스트
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstOfType(root *sitter.Node, typ string) *sitter.Node {
	nodes := findAllByType(root, typ)
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

func TestTryExtractDefault_Branches(t *testing.T) {
	// identifier / type / : / = nodes: from a typed default parameter
	src := []byte("def f(x: int = 5): pass\n")
	root, _ := parsePython(src)

	// ":" punctuation node
	if c := firstOfType(root, ":"); c != nil {
		v, call, none := tryExtractDefault(c, src)
		if v != "" || call != "" || none {
			t.Errorf(": node should yield empty, got %q %q %v", v, call, none)
		}
	}
	// identifier node ("x" or "int")
	if c := firstOfType(root, "identifier"); c != nil {
		v, call, none := tryExtractDefault(c, src)
		if v != "" || call != "" || none {
			t.Errorf("identifier should yield empty, got %q %q %v", v, call, none)
		}
	}
	// integer (default branch, non-None) -> returns text
	if c := firstOfType(root, "integer"); c != nil {
		v, _, none := tryExtractDefault(c, src)
		if v != "5" || none {
			t.Errorf("integer default: got %q none=%v", v, none)
		}
	}
}

func TestTryExtractDefault_None(t *testing.T) {
	src := []byte("def f(x: int = None): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "none")
	if c == nil {
		t.Skip("no none node found")
	}
	if _, _, none := tryExtractDefault(c, src); !none {
		t.Fatal("none node should set isNone")
	}
}

func TestTryExtractDefault_CallNoIdentifier(t *testing.T) {
	// call whose function is an attribute (obj.method) -> no top-level identifier
	src := []byte("def f(x = mod.factory()): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "call")
	if c == nil {
		t.Skip("no call node")
	}
	v, call, none := tryExtractDefault(c, src)
	if v == "" || none {
		t.Fatalf("call val should be non-empty: v=%q none=%v", v, none)
	}
	_ = call // may be "" when function is attribute access
}

func TestTryExtractDefault_Call(t *testing.T) {
	src := []byte("def f(x = Query(default=5)): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "call")
	if c == nil {
		t.Skip("no call node")
	}
	v, call, none := tryExtractDefault(c, src)
	if call != "Query" || v == "" || none {
		t.Fatalf("call: v=%q call=%q none=%v", v, call, none)
	}
}
