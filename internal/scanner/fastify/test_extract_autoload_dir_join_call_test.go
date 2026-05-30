//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractAutoloadDir_JoinCall 테스트
package fastify

import "testing"

func TestExtractAutoloadDir_JoinCall(t *testing.T) {
	obj, src := firstObject(t, `{ dir: join(__dirname, "routes", "api") }`)
	if got := extractAutoloadDir(obj, src); got != "routes/api" {
		t.Fatalf("join dir = %q, want routes/api", got)
	}
}
