//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestConvertOnePydanticField_Format 테스트
package fastapi

import "testing"

func TestConvertOnePydanticField_Format(t *testing.T) {

	sf := convertOnePydanticField(pydanticField{name: "created", typeName: "datetime"})
	if sf.Type == "" {
		t.Fatalf("empty type")
	}
}
