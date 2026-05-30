//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what mapJSONSchemaType 테스트
package fastify

import "testing"

func TestMapJSONSchemaType(t *testing.T) {
	tests := map[string]string{
		"string":  "string",
		"integer": "integer",
		"number":  "number",
		"boolean": "boolean",
		"array":   "array",
		"object":  "object",
		"custom":  "custom", // default: passthrough
	}
	for in, want := range tests {
		if got := mapJSONSchemaType(in); got != want {
			t.Errorf("mapJSONSchemaType(%q) = %q, want %q", in, got, want)
		}
	}
}
