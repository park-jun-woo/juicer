//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 엔드포인트 path/paths의 enum 멤버표현식을 일괄 해석한다
package nestjs

// resolveEndpointEnumPaths resolves enum member-expression method paths
// (e.g. @Get(RouteKey.X)) in place for each endpoint's path and paths slice.
func resolveEndpointEnumPaths(eps []endpointInfo, pc enumPathCtx) {
	for i := range eps {
		eps[i].path = pc.resolve(eps[i].path)
		for j := range eps[i].paths {
			eps[i].paths[j] = pc.resolve(eps[i].paths[j])
		}
	}
}
