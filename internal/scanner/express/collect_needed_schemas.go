//ff:func feature=scan type=extract control=iteration dimension=2 topic=express
//ff:what 모든 파일의 라우트에서 참조하는 스키마 변수명을 수집한다
package express

func collectNeededSchemas(ctx *scanContext) []string {
	var names []string
	seen := make(map[string]bool)
	for path, fi := range ctx.parsed {
		routers := ctx.allRouters[path]
		if len(routers) == 0 {
			continue
		}
		routes := extractRoutes(fi, routers)
		collectSchemaNames(routes, seen, &names)
	}
	return names
}
