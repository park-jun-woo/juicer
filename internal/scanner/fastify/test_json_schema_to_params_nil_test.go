//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToParams_Nil 테스트
package fastify

import "testing"

func TestJSONSchemaToParams_Nil(t *testing.T) {
	if got := jsonSchemaToParams(nil, []byte("")); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
