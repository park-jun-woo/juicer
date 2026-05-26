//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what X | None 패턴에서 nullable 타입을 추출한다
package fastapi

// tryPipeNullable handles X | None patterns.
func tryPipeNullable(py string) (openAPIType, bool) {
	parts := splitTopLevel(py)
	nonNone := filterNone(parts)
	if len(nonNone) == 1 && len(nonNone) < len(parts) {
		result := pyTypeToOpenAPI(nonNone[0])
		result.Nullable = true
		return result, true
	}
	return openAPIType{}, false
}
