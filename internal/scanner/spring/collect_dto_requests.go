//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 엔드포인트에서 DTO 해석 요청을 수집한다
package spring

func collectDTORequests(ep endpointInfo, ci controllerInfo, projectRoot string, epIdx int) []dtoRequest {
	var reqs []dtoRequest
	if ep.bodyType != "" && !isPrimitiveType(ep.bodyType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.bodyType,
			imports:     ci.imports,
			referrer:    ci.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
			isBody:      true,
		})
	}
	if ep.formType != "" && !isPrimitiveType(ep.formType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.formType,
			imports:     ci.imports,
			referrer:    ci.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
			isForm:      true,
		})
	}
	if ep.returnType != "" && !isPrimitiveType(ep.returnType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.returnType,
			imports:     ci.imports,
			referrer:    ci.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
		})
	}
	return reqs
}
