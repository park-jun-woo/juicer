//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what Union[X, None] 패턴에서 nullable 타입을 추출한다
package fastapi

// tryUnionNullable handles Union[X, None] patterns.
func tryUnionNullable(py string) (openAPIType, bool) {
	inner := py[6 : len(py)-1]
	parts := splitTopLevel(inner)
	nonNone := filterNone(parts)
	if len(nonNone) == 1 && len(nonNone) < len(parts) {
		result := pyTypeToOpenAPI(nonNone[0])
		result.Nullable = true
		return result, true
	}
	return openAPIType{}, false
}
