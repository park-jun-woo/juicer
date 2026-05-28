//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 어노테이션 없는 POJO 파라미터를 request body로 감지한다
package quarkus

func classifyBodyParam(typeName, paramName string, ep *endpointInfo) {
	if typeName == "" || primitiveTypes[typeName] {
		return
	}
	if ep.bodyType != "" {
		return
	}
	ep.bodyType = typeName
	ep.bodyVarName = paramName
}
