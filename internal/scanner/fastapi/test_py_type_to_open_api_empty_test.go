//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPyTypeToOpenAPI_Empty 테스트
package fastapi

import "testing"

func TestPyTypeToOpenAPI_Empty(t *testing.T) {
	oa := pyTypeToOpenAPI("")
	if oa.Type != "" {
		t.Fatalf("expected empty, got %s", oa.Type)
	}
}
