//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what HTTP 메서드 호출 패턴(path 형태 가드 통과)에서 추정 라우터 변수명을 수집한다 (타입 어노테이션 폴백)
package express

func collectParamRoutersFromUsage(fi *fileInfo, routers map[string]bool) {
	for _, call := range findAllByType(fi.Root, "call_expression") {
		if varName := usageRouterCandidate(call, fi.Src); varName != "" {
			routers[varName] = true
		}
	}
}
