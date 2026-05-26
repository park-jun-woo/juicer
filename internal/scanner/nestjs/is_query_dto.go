//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what @Query() 파라미터 타입이 DTO 전개 대상인지 확인한다
package nestjs

// isQueryDTO returns true if a @Query() decorator with empty arg has a DTO (object) type.
func isQueryDTO(arg, paramType string) bool {
	return arg == "" && tsTypeToOpenAPI(paramType).Type == "object"
}
