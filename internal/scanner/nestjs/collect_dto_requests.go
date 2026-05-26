//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 엔드포인트 정보에서 DTO 해석 요청을 수집한다
package nestjs

// collectDTORequests collects DTO resolution requests from an endpoint's body and return types.
func collectDTORequests(ep endpointInfo, imports map[string]string, absFile, projectRoot string, epIdx int) []dtoRequest {
	var reqs []dtoRequest
	if ep.bodyType != "" {
		reqs = append(reqs, dtoRequest{
			typeName: ep.bodyType, imports: imports,
			referrer: absFile, projectRoot: projectRoot,
			epIdx: epIdx, isBody: true,
		})
	}
	if ep.queryDTOType != "" {
		reqs = append(reqs, dtoRequest{
			typeName: ep.queryDTOType, imports: imports,
			referrer: absFile, projectRoot: projectRoot,
			epIdx: epIdx, isQuery: true,
		})
	}
	if needsResponseDTO(ep.returnType) {
		reqs = append(reqs, dtoRequest{
			typeName: ep.returnType, imports: imports,
			referrer: absFile, projectRoot: projectRoot,
			epIdx: epIdx, isBody: false,
		})
	}
	return reqs
}
