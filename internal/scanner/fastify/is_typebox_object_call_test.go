//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what isTypeBoxObjectCall 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstCall(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatalf("no call in %q", src)
	}
	return calls[0], fi.Src
}

func TestIsTypeBoxObjectCall_Match(t *testing.T) {
	c, src := firstCall(t, "const x = Type.Object({});\n")
	if !isTypeBoxObjectCall(c, src) {
		t.Error("expected Type.Object() to match")
	}
}

func TestIsTypeBoxObjectCall_NotObject(t *testing.T) {
	c, src := firstCall(t, "const x = Type.String();\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("Type.String() should not match")
	}
}

func TestIsTypeBoxObjectCall_NotType(t *testing.T) {
	c, src := firstCall(t, "const x = Other.Object({});\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("Other.Object() should not match")
	}
}

func TestIsTypeBoxObjectCall_DeepMember(t *testing.T) {
	// Type.Object as part of a longer chain still recognized at top
	c, src := firstCall(t, "const x = Type.Object({ a: Type.Number() });\n")
	if !isTypeBoxObjectCall(c, src) {
		t.Error("expected Type.Object with nested call to match")
	}
}

func TestIsTypeBoxObjectCall_NoMember(t *testing.T) {
	c, src := firstCall(t, "const x = foo();\n")
	if isTypeBoxObjectCall(c, src) {
		t.Error("plain call should not match")
	}
}
