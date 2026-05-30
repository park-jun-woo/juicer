//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadRequireName_NotRequire 테스트
package fastify

import "testing"

func TestAutoloadRequireName_NotRequire(t *testing.T) {
	d, src := firstDeclarator(t, `const x = foo("@fastify/autoload");`+"\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for non-require, got %q", got)
	}
}
