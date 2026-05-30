//ff:func feature=scan type=test control=sequence topic=scanner
//ff:what resolvePrimitiveSchema 테스트 (round5)
package scanner

import "testing"

func TestResolvePrimitiveSchema_Round5(t *testing.T) {
	// scalar primitive
	schema, base, slice := resolvePrimitiveSchema("int")
	if slice || base != "int" || schema == nil || schema["type"] != "integer" {
		t.Fatalf("int: %v %q %v", schema, base, slice)
	}

	// slice primitive => array wrapper
	schema, base, slice = resolvePrimitiveSchema("[]string")
	if !slice || base != "string" || schema == nil || schema["type"] != "array" {
		t.Fatalf("[]string: %v %q %v", schema, base, slice)
	}
	items, _ := schema["items"].(map[string]any)
	if items == nil || items["type"] != "string" {
		t.Fatalf("[]string items: %v", schema["items"])
	}

	// non-primitive => nil schema
	schema, base, slice = resolvePrimitiveSchema("[]User")
	if schema != nil || base != "User" || !slice {
		t.Fatalf("[]User: %v %q %v", schema, base, slice)
	}
}
