//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToFields_Nil 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_Nil(t *testing.T) {
	if got := jsonSchemaToFields(nil, []byte("")); got != nil {
		t.Fatalf("expected nil for nil node, got %v", got)
	}
}
