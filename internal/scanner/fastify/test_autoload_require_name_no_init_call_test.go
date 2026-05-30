//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAutoloadRequireName_NoInitCall 테스트
package fastify

import "testing"

func TestAutoloadRequireName_NoInitCall(t *testing.T) {
	d, src := firstDeclarator(t, "const x = 5;\n")
	if got := autoloadRequireName(d, src); got != "" {
		t.Fatalf("expected empty for no call, got %q", got)
	}
}
