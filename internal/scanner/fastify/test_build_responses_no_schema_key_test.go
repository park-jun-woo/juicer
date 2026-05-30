//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildResponses_NoSchemaKey 테스트
package fastify

import "testing"

func TestBuildResponses_NoSchemaKey(t *testing.T) {
	obj, src := firstObject(t, `{ config: {} }`)
	if r := buildResponses(routeInfo{Schema: obj}, src); r != nil {
		t.Fatalf("expected nil when schema info nil, got %v", r)
	}
}
