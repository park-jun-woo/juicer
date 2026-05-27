//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 엔드포인트 타입 필드에서 제네릭 파라미터를 실제 타입으로 치환한다
package nestjs

// substituteTypes replaces generic type parameter names with their concrete
// type arguments in the returnType, bodyType, and queryDTOType fields of each
// endpoint. The replacement is word-boundary based to avoid partial matches.
func substituteTypes(endpoints []endpointInfo, typeMap map[string]string) {
	for i := range endpoints {
		endpoints[i].returnType = substituteOne(endpoints[i].returnType, typeMap)
		endpoints[i].bodyType = substituteOne(endpoints[i].bodyType, typeMap)
		endpoints[i].queryDTOType = substituteOne(endpoints[i].queryDTOType, typeMap)
	}
}
