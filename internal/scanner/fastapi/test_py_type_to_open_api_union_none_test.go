//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPyTypeToOpenAPI_UnionNone 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_UnionNone(t *testing.T) {
	oa := pyTypeToOpenAPI("str | None")
	if oa.Type != "string" {
		t.Fatalf("expected string, got %s", oa.Type)
	}
	if !oa.Nullable {
		t.Fatal("expected nullable")
	}
}
