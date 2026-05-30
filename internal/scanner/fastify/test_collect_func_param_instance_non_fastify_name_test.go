//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFuncParamInstance_NonFastifyName 테스트
package fastify

import "testing"

func TestCollectFuncParamInstance_NonFastifyName(t *testing.T) {
	fn, src := firstFunc(t, "function plugin(opts) {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if len(instances) != 0 {
		t.Fatalf("expected no instances, got %v", instances)
	}
}
