//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFastifyVars_None 테스트
package fastify

import "testing"

func TestCollectFastifyVars_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	instances := map[string]bool{}
	collectFastifyVars(fi, instances)
	if len(instances) != 0 {
		t.Fatalf("expected none, got %v", instances)
	}
}
