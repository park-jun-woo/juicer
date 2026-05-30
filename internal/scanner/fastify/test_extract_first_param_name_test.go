//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractFirstParamName 테스트
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
