//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestJSONSchemaToParams_Object 테스트
package fastify

import "testing"

func TestJSONSchemaToParams_Object(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", properties: { page: { type: "integer" }, q: { type: "string" } } }`)
	params := jsonSchemaToParams(obj, src)
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	names := map[string]string{}
	for _, p := range params {
		names[p.Name] = p.Type
	}
	if names["page"] != "integer" || names["q"] != "string" {
		t.Fatalf("param types wrong: %v", names)
	}
}
