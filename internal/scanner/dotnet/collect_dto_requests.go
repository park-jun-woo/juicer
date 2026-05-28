//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 엔드포인트에서 DTO 해석 요청을 수집한다
package dotnet

func collectDTORequests(ep endpointInfo, ci controllerInfo, projectRoot string, epIdx int) []dtoRequest {
	var reqs []dtoRequest
	if ep.bodyType != "" && !isPrimitiveType(ep.bodyType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.bodyType,
			usings:      ci.usings,
			referrer:    ci.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
			isBody:      true,
		})
	}
	if ep.returnType != "" && !isPrimitiveType(ep.returnType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.returnType,
			usings:      ci.usings,
			referrer:    ci.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
		})
	}
	return reqs
}
