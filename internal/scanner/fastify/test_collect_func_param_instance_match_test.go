//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectFuncParamInstance_Match 테스트
package fastify

import "testing"

func TestCollectFuncParamInstance_Match(t *testing.T) {
	fn, src := firstFunc(t, "function plugin(fastify) {}\n")
	instances := map[string]bool{}
	collectFuncParamInstance(fn, src, instances)
	if !instances["fastify"] {
		t.Fatalf("expected fastify instance, got %v", instances)
	}
}
