//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what convertOnePydanticField: 기본 string / format 부착 / required / default 분기
package fastapi

import "testing"

func TestConvertOnePydanticField_RequiredString(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "name", typeName: "str"})
	if sf.Name != "name" || sf.Type != "string" || sf.Validate != "required" {
		t.Fatalf("got %+v", sf)
	}
}

func TestConvertOnePydanticField_WithDefault(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "n", typeName: "int", hasDefault: true})
	if sf.Validate == "required" {
		t.Fatalf("expected not required, got %+v", sf)
	}
}

func TestConvertOnePydanticField_Format(t *testing.T) {
	// a type with an OpenAPI format (e.g. datetime -> string:date-time)
	sf := convertOnePydanticField(pydanticField{name: "created", typeName: "datetime"})
	if sf.Type == "" {
		t.Fatalf("empty type")
	}
}

func TestConvertOnePydanticField_UnknownTypeDefaultsString(t *testing.T) {
	sf := convertOnePydanticField(pydanticField{name: "x", typeName: "SomeUnknownCustomType"})
	if sf.Type == "" {
		t.Fatalf("expected non-empty type for unknown, got %+v", sf)
	}
}

func TestConvertOnePydanticField_EmptyTypeDefaultsString(t *testing.T) {
	// empty typeName -> pyTypeToOpenAPI yields empty Type -> defaults to "string"
	sf := convertOnePydanticField(pydanticField{name: "x", typeName: ""})
	if sf.Type != "string" {
		t.Fatalf("expected string default, got %+v", sf)
	}
}
