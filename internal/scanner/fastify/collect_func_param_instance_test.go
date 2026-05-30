//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what collectFuncParamInstance 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstFunc(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	for _, typ := range []string{"function_declaration", "function", "arrow_function"} {
		if ns := findAllByType(fi.Root, typ); len(ns) > 0 {
			return ns[0], fi.Src
		}
	}
	t.Fatalf("no function in %q", src)
	return nil, nil
}

func TestCollectFuncParamInstance_Match(t *testing.T) {
	fn, src := firstFunc(t, "function plugin(fastify) {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if !instances["fastify"] {
		t.Fatalf("expected fastify instance, got %v", instances)
	}
}

func TestCollectFuncParamInstance_NonFastifyName(t *testing.T) {
	fn, src := firstFunc(t, "function plugin(opts) {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances, got %v", instances)
	}
}

func TestCollectFuncParamInstance_NoParamName(t *testing.T) {
	// function with no parameters -> empty name
	fn, src := firstFunc(t, "function plugin() {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances for no-param func, got %v", instances)
	}
}

func TestCollectFuncParamInstance_NoFormalParameters(t *testing.T) {
	// arrow function with a single bare-identifier param has no formal_parameters
	// node -> params == nil early return.
	fi := mustParse(t, []byte("const f = fastify => fastify.get();\n"))
	arrows := findAllByType(fi.Root, "arrow_function")
	if len(arrows) == 0 {
		t.Fatal("no arrow_function")
	}
	instances := map[string]bool{}
	collectFuncParamInstance(arrows[0], fi.Src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances (no formal_parameters), got %v", instances)
	}
}
