//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what app.route("/prefix", subApp) 그룹 호출을 수집한다
package hono

func collectRouteGroups(fi *fileInfo, honoVars map[string]bool) []routeGroup {
	var groups []routeGroup
	calls := findAllByType(fi.Root, "call_expression")
	for _, call := range calls {
		g := extractRouteGroup(call, fi.Src, honoVars)
		if g != nil {
			groups = append(groups, *g)
		}
	}
	return groups
}
