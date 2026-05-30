//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what jsonSchemaToParams 테스트
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

func TestJSONSchemaToParams_Nil(t *testing.T) {
	if got := jsonSchemaToParams(nil, []byte("")); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestJSONSchemaToParams_NonObject(t *testing.T) {
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	if got := jsonSchemaToParams(n, src); got != nil {
		t.Fatalf("expected nil for non-object, got %v", got)
	}
}

func TestJSONSchemaToParams_NoProperties(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object" }`)
	if got := jsonSchemaToParams(obj, src); got != nil {
		t.Fatalf("expected nil when no properties, got %v", got)
	}
}

func TestJSONSchemaToParams_PropertiesNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ properties: 5 }`)
	if got := jsonSchemaToParams(obj, src); got != nil {
		t.Fatalf("expected nil when properties not object, got %v", got)
	}
}
