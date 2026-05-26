//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mapTypeToOpenAPI 테스트
package fastapi

import "testing"

func TestMapTypeToOpenAPI(t *testing.T) {
	if got := mapTypeToOpenAPI("str"); got != "string" {
		t.Errorf("str: got %s", got)
	}
	if got := mapTypeToOpenAPI("int"); got != "integer" {
		t.Errorf("int: got %s", got)
	}
	// Unknown uppercase type maps to "object" (Pydantic model)
	if got := mapTypeToOpenAPI("UnknownModel"); got == "" {
		t.Errorf("UnknownModel: got empty")
	}
}
