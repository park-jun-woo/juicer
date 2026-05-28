//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what app.route() 그룹 + basePath에서 prefix를 전파하여 varName→prefix 매핑을 생성한다
package hono

func resolveRoutePrefixes(groups []routeGroup, basePaths map[string]string) map[string]string {
	prefixMap := make(map[string]string)
	for varName, bp := range basePaths {
		prefixMap[varName] = bp
	}
	for _, g := range groups {
		parentPrefix := prefixMap[g.ParentVar]
		fullPrefix := joinHonoPath(parentPrefix, g.Prefix)
		prefixMap[g.SubAppName] = fullPrefix
	}
	return prefixMap
}
