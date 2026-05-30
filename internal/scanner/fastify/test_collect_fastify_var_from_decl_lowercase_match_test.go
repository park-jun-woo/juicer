//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestCollectFastifyVarFromDecl_LowercaseMatch 테스트
package fastify

import "testing"

func TestCollectFastifyVarFromDecl_LowercaseMatch(t *testing.T) {
	decls, fi := declOfType(t, "const srv = fastify();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["srv"] {
		t.Fatalf("expected 'srv' instance, got %v", instances)
	}
}
