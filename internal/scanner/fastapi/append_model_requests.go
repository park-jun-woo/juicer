//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what body 및 response 타입에 대한 모델 해석 요청을 추가한다
package fastapi

// appendModelRequests adds model resolution requests for body and response types.
func appendModelRequests(reqs []modelRequest, ri routeInfo, fi fileInfo, epIdx int) []modelRequest {
	if ri.bodyType != "" {
		reqs = append(reqs, modelRequest{
			typeName: ri.bodyType,
			imports:  fi.imports,
			referrer: fi.absPath,
			epIdx:    epIdx,
			isBody:   true,
		})
	}
	respType := ri.responseModel
	if respType == "" {
		respType = ri.returnType
	}
	if respType != "" && isPydanticModelType(respType) {
		reqs = append(reqs, modelRequest{
			typeName: respType,
			imports:  fi.imports,
			referrer: fi.absPath,
			epIdx:    epIdx,
			isBody:   false,
		})
	}
	return reqs
}
