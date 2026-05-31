//ff:func feature=scan type=test control=sequence
//ff:what TestResolvePrimitiveSchema_TSArraySuffix — TS식 T[] 접미 배열 처리
package scanner

import "testing"

func TestResolvePrimitiveSchema_TSArraySuffix(t *testing.T) {
	// Named TS array type must be treated as a slice with the [] stripped, so
	// the emit layer produces type:array + items:$ref, not a "xxx[]" pseudo-schema.
	schema, base, isSlice := resolvePrimitiveSchema("AssetResponseDto[]")
	if schema != nil {
		t.Fatalf("named type must not resolve to a primitive schema, got %v", schema)
	}
	if base != "AssetResponseDto" {
		t.Fatalf("baseName: want %q got %q", "AssetResponseDto", base)
	}
	if !isSlice {
		t.Fatal("expected isSlice=true for T[] suffix")
	}
}
