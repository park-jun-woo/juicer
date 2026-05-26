//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 반환 타입이 DTO 해석이 필요한 객체 타입인지 확인한다
package nestjs

// needsResponseDTO checks if a return type needs DTO resolution.
func needsResponseDTO(returnType string) bool {
	if returnType == "" || returnType == "any" {
		return false
	}
	oa := tsTypeToOpenAPI(returnType)
	return oa.Type == "object"
}
