//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParamNodeName_RequiredParameter 테스트
package fastify

import "testing"

func TestParamNodeName_RequiredParameter(t *testing.T) {
	fi := mustParse(t, []byte("function f(a: number) {}\n"))
	params := findAllByType(fi.Root, "formal_parameters")[0]
	rp := findChildByType(params, "required_parameter")
	if rp == nil {
		t.Skip("no required_parameter")
	}
	if got := paramNodeName(rp, fi.Src); got != "a" {
		t.Fatalf("required_parameter: got %q", got)
	}
}
