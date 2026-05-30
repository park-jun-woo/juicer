//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParamNodeName_Other 테스트
package fastify

import "testing"

func TestParamNodeName_Other(t *testing.T) {

	fi := mustParse(t, []byte("function f(a) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	if got := paramNodeName(params.Child(0), fi.Src); got != "" {
		t.Fatalf("expected empty for '(' node, got %q", got)
	}
}
