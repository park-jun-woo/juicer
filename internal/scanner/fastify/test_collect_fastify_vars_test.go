//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFastifyVars 테스트
package fastify

import "testing"

func TestCollectFastifyVars(t *testing.T) {
	src := `
const app = Fastify();
const router = fastify();
const x = 5;
`
	fi := mustParse(t, []byte(src))
	instances := map[string]bool{}
	collectFastifyVars(fi, instances)
	if !instances["app"] || !instances["router"] {
		t.Fatalf("expected app and router, got %v", instances)
	}
	if instances["x"] {
		t.Fatalf("x should not be an instance, got %v", instances)
	}
}
