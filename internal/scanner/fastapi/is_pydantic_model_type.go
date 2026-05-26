//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 타입명이 Pydantic 모델인지 휴리스틱으로 판별한다
package fastapi

// isPydanticModelType heuristically detects if a type name is a Pydantic model.
// Model types start with uppercase and are not Python builtins.
func isPydanticModelType(typeName string) bool {
	if typeName == "" {
		return false
	}
	switch typeName {
	case "str", "int", "float", "bool", "list", "dict", "set", "tuple",
		"List", "Dict", "Set", "Tuple", "Optional", "Union", "Any",
		"datetime", "date", "EmailStr", "UploadFile", "None":
		return false
	}
	first := typeName[0]
	return first >= 'A' && first <= 'Z'
}
