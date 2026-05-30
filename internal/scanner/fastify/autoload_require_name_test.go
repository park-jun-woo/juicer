//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what autoloadRequireName 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstDeclarator(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ds := findAllByType(fi.Root, "variable_declarator")
	if len(ds) == 0 {
		t.Fatalf("no variable_declarator in %q", src)
	}
	return ds[0], fi.Src
}

func TestAutoloadRequireName_Match(t *testing.T) {
	d, src := firstDeclarator(t, `const autoload = require("@fastify/autoload");`+"\n")
	if got := autoloadRequireName(d, src); got != "autoload" {
		t.Fatalf("expected autoload, got %q", got)
	}
}

func TestAutoloadRequireName_NoInitCall(t *testing.T) {
	d, src := firstDeclarator(t, "const x = 5;\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for no call, got %q", got)
	}
}

func TestAutoloadRequireName_NotRequire(t *testing.T) {
	d, src := firstDeclarator(t, `const x = foo("@fastify/autoload");`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for non-require, got %q", got)
	}
}

func TestAutoloadRequireName_WrongModule(t *testing.T) {
	d, src := firstDeclarator(t, `const x = require("other-module");`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for wrong module, got %q", got)
	}
}

func TestAutoloadRequireName_RequireNoArg(t *testing.T) {
	// require() with no string arg -> "" (extractCallStringArg mismatch)
	d, src := firstDeclarator(t, `const x = require();`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for require with no arg, got %q", got)
	}
}
