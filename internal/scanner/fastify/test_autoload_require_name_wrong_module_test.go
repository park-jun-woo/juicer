//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadRequireName_WrongModule 테스트
package fastify

import "testing"

func TestAutoloadRequireName_WrongModule(t *testing.T) {
	d, src := firstDeclarator(t, `const x = require("other-module");`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for wrong module, got %q", got)
	}
}
