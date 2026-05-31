//ff:func feature=scan type=test control=sequence
//ff:what TestResolvePrimitiveSchema_TSArraySuffix_Primitive — string[] 인라인 배열
package scanner

import "testing"

func TestResolvePrimitiveSchema_TSArraySuffix_Primitive(t *testing.T) {
	// string[] -> array of strings inline.
	schema, base, isSlice := resolvePrimitiveSchema("string[]")
	if !isSlice || base != "string" {
		t.Fatalf("expected slice of string, got base=%q isSlice=%v", base, isSlice)
	}
	if schema == nil || schema["type"] != "array" {
		t.Fatalf("expected inline array schema, got %v", schema)
	}
}
