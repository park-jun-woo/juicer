//ff:func feature=scan type=extract control=sequence
//ff:what []접두사를 분리하고 primitive이면 inline schema(배열 포함)를 반환한다
package scanner

import "strings"

// resolvePrimitiveSchema strips the [] prefix from rawTypeName, looks up
// primitiveTypeSchema, and returns the resolved schema (array-wrapped if
// slice) or nil when the type is not primitive.
func resolvePrimitiveSchema(rawTypeName string) (schema map[string]any, baseName string, isSlice bool) {
	isSlice = strings.HasPrefix(rawTypeName, "[]")
	baseName = rawTypeName
	if isSlice {
		baseName = rawTypeName[2:]
	}

	prim, ok := primitiveTypeSchema[baseName]
	if !ok {
		return nil, baseName, isSlice
	}

	if isSlice {
		return map[string]any{"type": "array", "items": prim}, baseName, isSlice
	}
	return prim, baseName, isSlice
}
