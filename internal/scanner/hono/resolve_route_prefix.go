//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what app.route() 그룹 + basePath에서 prefix를 파일-aware하게 전파하여 (파일,변수)→prefix 매핑을 생성한다
package hono

func resolveRoutePrefixes(groups []routeGroup, basePaths map[string]string, honoVars map[string]map[string]bool, imports map[string]map[string]string) map[string]string {
	prefixMap := make(map[string]string)
	for key, bp := range basePaths {
		prefixMap[key] = bp
	}
	for i := 0; i <= len(groups); i++ {
		if !runPrefixPass(groups, prefixMap, honoVars, imports) {
			break
		}
	}
	return prefixMap
}
