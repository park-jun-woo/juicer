//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what Fastify() 호출 인스턴스 변수 수집 테스트
package fastify

import "testing"

func TestCollectInstances_Fastify(t *testing.T) {
	src := []byte(`
import Fastify from "fastify";
const app = Fastify();
`)
	fi := mustParse(t, src)
	instances := collectInstances(fi)
	if !instances["app"] {
		t.Error("expected 'app' in instances")
	}
}
