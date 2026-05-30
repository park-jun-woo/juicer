//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractFirstParamName_Empty 테스트
package fastify

import "testing"

func TestExtractFirstParamName_Empty(t *testing.T) {
	fi := mustParse(t, []byte("function f() {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := extractFirstParamName(params, fi.Src); got != "" {
		t.Fatalf("expected empty for no params, got %q", got)
	}
}
