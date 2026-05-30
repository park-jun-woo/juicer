//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestConvertOnePydanticField_RequiredString 테스트
package fastapi

import "testing"

func TestConvertOnePydanticField_RequiredString(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "name", typeName: "str"})
	if sf.Name != "name" || sf.Type != "string" || sf.Validate != "required" {
		t.Fatalf("got %+v", sf)
	}
}
