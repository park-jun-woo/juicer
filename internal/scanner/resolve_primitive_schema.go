//ff:func feature=scan type=extract control=selection
//ff:what 배열 마커([]접두/접미)를 분리하고 primitive이면 inline schema(배열 포함)를 반환한다
package scanner

import "strings"

// resolvePrimitiveSchema strips the array marker from rawTypeName (Go-style
// "[]T" prefix or TypeScript-style "T[]" suffix), looks up primitiveTypeSchema,
// and returns the resolved schema (array-wrapped if slice) or nil when the type
// is not primitive. The suffix form lets named TS array types such as
// "XxxResponseDto[]" resolve to a proper type:array + items:$ref instead of a
// pseudo-schema named "xxxResponseDto[]".
func resolvePrimitiveSchema(rawTypeName string) (schema map[string]any, baseName string, isSlice bool) {
	baseName = rawTypeName
	switch {
	case strings.HasPrefix(rawTypeName, "[]"):
		isSlice = true
		baseName = rawTypeName[2:]
	case strings.HasSuffix(rawTypeName, "[]"):
		isSlice = true
		baseName = rawTypeName[:len(rawTypeName)-2]
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
