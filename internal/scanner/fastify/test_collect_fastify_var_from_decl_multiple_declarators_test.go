//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestCollectFastifyVarFromDecl_MultipleDeclarators 테스트
package fastify

import "testing"

func TestCollectFastifyVarFromDecl_MultipleDeclarators(t *testing.T) {

	decls, fi := declOfType(t, "const app = Fastify(), other = 5;\n")
	instances := map[string]bool{}
	for _, d := range decls {
		collectFastifyVarFromDecl(d, fi, instances)
	}
	if !instances["app"] || instances["other"] {
		t.Fatalf("expected only app, got %v", instances)
	}
}
