//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 엔드포인트에서 DTO 해석 요청을 수집한다
package quarkus

func collectDTORequests(ep endpointInfo, ri resourceInfo, projectRoot string, epIdx int) []dtoRequest {
	var reqs []dtoRequest
	if ep.bodyType != "" && !isPrimitiveType(ep.bodyType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.bodyType,
			imports:     ri.imports,
			referrer:    ri.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
			isBody:      true,
		})
	}
	if ep.formType != "" && !isPrimitiveType(ep.formType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.formType,
			imports:     ri.imports,
			referrer:    ri.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
			isForm:      true,
		})
	}
	if ep.returnType != "" && !isPrimitiveType(ep.returnType) {
		reqs = append(reqs, dtoRequest{
			typeName:    ep.returnType,
			imports:     ri.imports,
			referrer:    ri.absFile,
			projectRoot: projectRoot,
			epIdx:       epIdx,
		})
	}
	return reqs
}
