//ff:func feature=scan type=extract control=sequence
//ff:what primitive type name을 OpenAPI inline schema로 매핑한다
package scanner

// primitiveTypeSchema maps language-level primitive type names to their
// OpenAPI inline schema. Used by responseSchema and bodySchema to avoid
// generating $ref for primitives.
var primitiveTypeSchema = map[string]map[string]any{
	// Python builtins
	"bool":  {"type": "boolean"},
	"int":   {"type": "integer"},
	"float": {"type": "number"},
	"str":   {"type": "string"},
	// TypeScript builtins
	"boolean": {"type": "boolean"},
	"number":  {"type": "number"},
	"string":  {"type": "string"},
	// 공통
	"any":  {"type": "object"},
	"dict": {"type": "object"},
}

func isPrimitiveTypeName(typeName string) bool {
	_, ok := primitiveTypeSchema[typeName]
	return ok
}
