//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPyTypeToOpenAPI_Optional 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_Optional(t *testing.T) {
	oa := pyTypeToOpenAPI("Optional[str]")
	if oa.Type != "string" {
		t.Fatalf("expected string, got %s", oa.Type)
	}
	if !oa.Nullable {
		t.Fatal("expected nullable")
	}
}
