//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFuncParamInstance_NoFormalParameters 테스트
package fastify

import "testing"

func TestCollectFuncParamInstance_NoFormalParameters(t *testing.T) {

	fi := mustParse(t, []byte("const f = fastify => fastify.get();\n"))
	arrows := findAllByType(fi.Root, "arrow_function")
	if len(arrows) == 0 {
		t.Fatal("no arrow_function")
	}
	instances := map[string]bool{}
	collectFuncParamInstance(arrows[0], fi.Src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances (no formal_parameters), got %v", instances)
	}
}
