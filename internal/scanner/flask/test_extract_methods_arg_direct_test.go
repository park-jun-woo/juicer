//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractMethodsArg_Direct 테스트
package flask

import "testing"

func TestExtractMethodsArg_Direct(t *testing.T) {
	args, src := argListOf(t, `route('/x', methods=['GET', 'POST', 'PUT'])`+"\n")
	got := extractMethodsArg(args, src)
	if len(got) != 3 || got[0] != "GET" || got[2] != "PUT" {
		t.Fatalf("methods = %v", got)
	}
}
