//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestCollectFastifyVarFromDecl_NonInstances 테스트
package fastify

import "testing"

func TestCollectFastifyVarFromDecl_NonInstances(t *testing.T) {

	decls, fi := declOfType(t, "const a = 5;\nconst b = obj.create();\nconst c = Express();\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if len(instances) != 0 {
		t.Fatalf("expected no instances, got %v", instances)
	}
}
