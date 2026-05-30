//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestFirstIdentArg 테스트
package flask

import "testing"

func TestFirstIdentArg(t *testing.T) {
	args, src := argListOf(t, `register_blueprint(api, url_prefix='/x')`+"\n")
	if got := firstIdentArg(args, src); got != "api" {
		t.Fatalf("got %q, want api", got)
	}
}
