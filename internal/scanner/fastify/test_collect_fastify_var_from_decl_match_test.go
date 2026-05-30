//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestCollectFastifyVarFromDecl_Match 테스트
package fastify

import "testing"

func TestCollectFastifyVarFromDecl_Match(t *testing.T) {
	decls, fi := declOfType(t, "const app = Fastify();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["app"] {
		t.Fatalf("expected 'app' instance, got %v", instances)
	}
}
