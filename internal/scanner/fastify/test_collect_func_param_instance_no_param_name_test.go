//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFuncParamInstance_NoParamName 테스트
package fastify

import "testing"

func TestCollectFuncParamInstance_NoParamName(t *testing.T) {

	fn, src := firstFunc(t, "function plugin() {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances for no-param func, got %v", instances)
	}
}
