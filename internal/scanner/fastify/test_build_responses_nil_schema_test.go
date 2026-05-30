//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildResponses_NilSchema 테스트
package fastify

import "testing"

func TestBuildResponses_NilSchema(t *testing.T) {
	if r := buildResponses(routeInfo{}, []byte("")); r != nil {
		t.Fatalf("expected nil for nil schema, got %v", r)
	}
}
