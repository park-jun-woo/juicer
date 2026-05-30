//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFileAutoloadMounts_NoNames 테스트
package fastify

import "testing"

func TestFileAutoloadMounts_NoNames(t *testing.T) {

	fi := mustParse(t, []byte("const app = Fastify();\n"))
	if m := fileAutoloadMounts("/app/server.ts", fi, "/app"); m != nil {
		t.Fatalf("expected nil when no autoload names, got %v", m)
	}
}
