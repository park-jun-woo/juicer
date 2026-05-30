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
	// UUID preserves format via "type:format" convention
	if got := mapTypeToOpenAPI("uuid.UUID"); got != "string:uuid" {
		t.Errorf("uuid.UUID: got %s, want string:uuid", got)
	}
	// datetime preserves format
	if got := mapTypeToOpenAPI("datetime"); got != "string:date-time" {
		t.Errorf("datetime: got %s, want string:date-time", got)
	}
	// empty type -> defaults to string
	if got := mapTypeToOpenAPI(""); got != "string" {
		t.Errorf("empty: got %s, want string", got)
	}
}
