//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestConvertOnePydanticField_UnknownTypeDefaultsString 테스트
package fastapi

import "testing"

func TestConvertOnePydanticField_UnknownTypeDefaultsString(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "x", typeName: "SomeUnknownCustomType"})
	if sf.Type == "" {
		t.Fatalf("expected non-empty type for unknown, got %+v", sf)
	}
}
