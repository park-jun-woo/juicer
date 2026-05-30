//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestConvertOnePydanticField_WithDefault 테스트
package fastapi

import "testing"

func TestConvertOnePydanticField_WithDefault(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "n", typeName: "int", hasDefault: true})
	if sf.Validate == "required" {
		t.Fatalf("expected not required, got %+v", sf)
	}
}
