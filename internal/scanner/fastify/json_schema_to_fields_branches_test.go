//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what jsonSchemaToFields 분기 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_Nil(t *testing.T) {
	if got := jsonSchemaToFields(nil, []byte("")); got != nil {
		t.Fatalf("expected nil for nil node, got %v", got)
	}
}

func TestJSONSchemaToFields_NonObject(t *testing.T) {
	// a string node, not an object -> nil
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	if got := jsonSchemaToFields(n, src); got != nil {
		t.Fatalf("expected nil for non-object, got %v", got)
	}
}

func TestJSONSchemaToFields_Identifier(t *testing.T) {
	// a TypeBox variable reference (identifier) -> nil
	fi := mustParse(t, []byte("const x = SomeSchema;\n"))
	ids := findAllByType(fi.Root, "identifier")
	var ref = ids[len(ids)-1] // the RHS reference
	if got := jsonSchemaToFields(ref, fi.Src); got != nil {
		t.Fatalf("expected nil for identifier, got %v", got)
	}
}

func TestJSONSchemaToFields_NoProperties(t *testing.T) {
	// object with type but no properties key -> nil
	obj, src := firstObject(t, `{ type: "object" }`)
	if got := jsonSchemaToFields(obj, src); got != nil {
		t.Fatalf("expected nil when no properties, got %v", got)
	}
}

func TestJSONSchemaToFields_PropertiesNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", properties: "x" }`)
	if got := jsonSchemaToFields(obj, src); got != nil {
		t.Fatalf("expected nil when properties not object, got %v", got)
	}
}
