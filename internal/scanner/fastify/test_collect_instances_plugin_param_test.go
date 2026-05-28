//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what 플러그인 함수 파라미터 인스턴스 수집 테스트
package fastify

import "testing"

func TestCollectInstances_PluginParam(t *testing.T) {
	src := []byte(`
export default async function(fastify) {
  fastify.get("/items", listItems);
}
`)
	fi := mustParse(t, src)
	instances := collectInstances(fi)
	if !instances["fastify"] {
		t.Error("expected 'fastify' in instances")
	}
}
