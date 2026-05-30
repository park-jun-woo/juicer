//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadRequireName_RequireNoArg 테스트
package fastify

import "testing"

func TestAutoloadRequireName_RequireNoArg(t *testing.T) {

	d, src := firstDeclarator(t, `const x = require();`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for require with no arg, got %q", got)
	}
}
