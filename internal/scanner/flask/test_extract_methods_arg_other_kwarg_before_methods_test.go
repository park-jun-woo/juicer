//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractMethodsArg_OtherKwargBeforeMethods 테스트
package flask

import "testing"

func TestExtractMethodsArg_OtherKwargBeforeMethods(t *testing.T) {

	args, src := argListOf(t, `route('/x', strict_slashes=False, methods=['GET'])`+"\n")
	got := extractMethodsArg(args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("methods = %v", got)
	}
}
