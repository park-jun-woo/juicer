//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestFirstStringArg_None 테스트
package flask

import "testing"

func TestFirstStringArg_None(t *testing.T) {
	args, src := argListOf(t, `f(x, y)`+"\n")
	if got := firstStringArg(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
