//ff:func feature=scan type=test control=sequence topic=flask
//ff:what firstStringArg 테스트
package flask

import "testing"

func TestFirstStringArg(t *testing.T) {
	args, src := argListOf(t, `route('/users', methods=['GET'])`+"\n")
	if got := firstStringArg(args, src); got != "/users" {
		t.Fatalf("got %q, want /users", got)
	}
}

func TestFirstStringArg_None(t *testing.T) {
	args, src := argListOf(t, `f(x, y)`+"\n")
	if got := firstStringArg(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
