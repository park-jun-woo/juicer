//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractPluginRef 테스트
package fastify

import "testing"

func TestExtractPluginRef_Identifier(t *testing.T) {
	fi := mustParse(t, []byte("foo(myPlugin);\n"))
	// the argument identifier inside the call
	args := findAllByType(fi.Root, "arguments")[0]
	id := findChildByType(args, "identifier")
	if id == nil {
		t.Fatal("no identifier arg")
	}
	if got := extractPluginRef(id, fi.Src); got != "myPlugin" {
		t.Fatalf("identifier: got %q", got)
	}
}

func TestExtractPluginRef_CallExpression(t *testing.T) {
	n, src := firstNodeOfType(t, `const x = require("@fastify/cors");`+"\n", "call_expression")
	if got := extractPluginRef(n, src); got != "@fastify/cors" {
		t.Fatalf("call: got %q", got)
	}
}

func TestExtractPluginRef_Arrow(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = () => 1;\n", "arrow_function")
	if got := extractPluginRef(n, src); got != inlineRef {
		t.Fatalf("arrow: got %q, want %q", got, inlineRef)
	}
}

func TestExtractPluginRef_Default(t *testing.T) {
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	if got := extractPluginRef(n, src); got != "" {
		t.Fatalf("default: got %q, want empty", got)
	}
}
