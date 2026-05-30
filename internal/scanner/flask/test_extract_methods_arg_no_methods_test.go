//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractMethodsArg_NoMethods 테스트
package flask

import "testing"

func TestExtractMethodsArg_NoMethods(t *testing.T) {
	args, src := argListOf(t, `route('/x')`+"\n")
	if got := extractMethodsArg(args, src); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
