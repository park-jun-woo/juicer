//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestConvertOnePydanticField_EmptyTypeDefaultsString 테스트
package fastapi

import "testing"

func TestConvertOnePydanticField_EmptyTypeDefaultsString(t *testing.T) {

	sf := convertOnePydanticField(pydanticField{name: "x", typeName: ""})
	if sf.Type != "string" {
		t.Fatalf("expected string default, got %+v", sf)
	}
}
