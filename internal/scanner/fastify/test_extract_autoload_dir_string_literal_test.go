//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadDir_StringLiteral 테스트
package fastify

import "testing"

func TestExtractAutoloadDir_StringLiteral(t *testing.T) {
	obj, src := firstObject(t, `{ dir: "routes" }`)
	if got := extractAutoloadDir(obj, src); got != "routes" {
		t.Fatalf("string dir = %q", got)
	}
}
