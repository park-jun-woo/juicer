//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestFirstIdentArg_None 테스트
package flask

import "testing"

func TestFirstIdentArg_None(t *testing.T) {
	args, src := argListOf(t, `f("only_string")`+"\n")
	if got := firstIdentArg(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
