//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractJSONSchema_NilOpts 테스트
package fastify

import "testing"

func TestExtractJSONSchema_NilOpts(t *testing.T) {
	if si := extractJSONSchema(nil, []byte("")); si != nil {
		t.Fatalf("expected nil for nil opts, got %v", si)
	}
}
