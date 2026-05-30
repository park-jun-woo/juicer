//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToFields_NonObject 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_NonObject(t *testing.T) {

	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	if got := jsonSchemaToFields(n, src); got != nil {
		t.Fatalf("expected nil for non-object, got %v", got)
	}
}
