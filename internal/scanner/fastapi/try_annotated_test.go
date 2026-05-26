//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryAnnotated 테스트
package fastapi

import "testing"

func TestTryAnnotated(t *testing.T) {
	// Simple int
	oa, ok := tryAnnotated("Annotated[int, Field(ge=0)]")
	if !ok || oa.Type != "integer" {
		t.Fatalf("Annotated[int]: got %v, %v", oa, ok)
	}
	// String
	oa2, ok := tryAnnotated("Annotated[str, Query()]")
	if !ok || oa2.Type != "string" {
		t.Fatalf("Annotated[str]: got %v, %v", oa2, ok)
	}
	// Nested list
	oa3, ok := tryAnnotated("Annotated[list[int], Field()]")
	if !ok || oa3.Type != "array" || oa3.Items != "int" {
		t.Fatalf("Annotated[list[int]]: got %v, %v", oa3, ok)
	}
	// UUID
	oa4, ok := tryAnnotated("Annotated[uuid.UUID, Path()]")
	if !ok || oa4.Type != "string" || oa4.Format != "uuid" {
		t.Fatalf("Annotated[uuid.UUID]: got %v, %v", oa4, ok)
	}
	// Not Annotated
	_, ok = tryAnnotated("str")
	if ok {
		t.Fatal("should not match plain type")
	}
}
