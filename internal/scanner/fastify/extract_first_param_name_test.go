//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractFirstParamName 테스트
package fastify

import "testing"

func TestExtractFirstParamName(t *testing.T) {
	fi := mustParse(t, []byte("function plugin(fastify, opts) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")
	if len(params) == 0 {
		t.Fatal("no formal_parameters")
	}
	if got := extractFirstParamName(params[0], fi.Src); got != "fastify" {
		t.Fatalf("expected fastify, got %q", got)
	}
}

func TestExtractFirstParamName_Empty(t *testing.T) {
	fi := mustParse(t, []byte("function f() {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := extractFirstParamName(params, fi.Src); got != "" {
		t.Fatalf("expected empty for no params, got %q", got)
	}
}

func TestExtractFirstParamName_Typed(t *testing.T) {
	// required_parameter with a type annotation
	fi := mustParse(t, []byte("function f(app: FastifyInstance) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := extractFirstParamName(params, fi.Src); got != "app" {
		t.Fatalf("expected app, got %q", got)
	}
}
