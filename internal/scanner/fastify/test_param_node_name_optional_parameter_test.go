//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParamNodeName_OptionalParameter 테스트
package fastify

import "testing"

func TestParamNodeName_OptionalParameter(t *testing.T) {
	fi := mustParse(t, []byte("function f(a?: number) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	op := findChildByType(params, "optional_parameter")
	if op == nil {
		t.Skip("no optional_parameter")
	}
	if got := paramNodeName(op, fi.Src); got != "a" {
		t.Fatalf("optional_parameter: got %q", got)
	}
}
