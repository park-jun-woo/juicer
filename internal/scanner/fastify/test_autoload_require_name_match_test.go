//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadRequireName_Match 테스트
package fastify

import "testing"

func TestAutoloadRequireName_Match(t *testing.T) {
	d, src := firstDeclarator(t, `const autoload = require("@fastify/autoload");`+"\n")
	if got := autoloadRequireName(d, src); got != "autoload" {
		t.Fatalf("expected autoload, got %q", got)
	}
}
