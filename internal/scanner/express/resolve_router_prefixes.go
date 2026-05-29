//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 마운트 그래프를 풀어 각 라우터 인스턴스(file,var)의 전체 prefix 목록을 계산한다
package express

// resolveRouterPrefixes는 모든 마운트 엔트리와 라우터 집합으로부터
// 라우터 인스턴스별 전체 prefix 목록(map[routerKey][]string)을 계산한다.
//
// 파일 단위가 아니라 (file, varName) 단위로 키잉하므로:
//   - 같은 라우터를 두 prefix에 마운트하면 두 prefix가 모두 보존된다.
//   - 한 파일에 라우터가 여럿이면 각각 다른 prefix를 받는다.
func resolveRouterPrefixes(mounts []mountEntry, allRouters map[string]map[string]bool) map[routerKey][]string {
	g := buildMountGraph(mounts, allRouters)
	prefixes := initRootPrefixes(g)
	propagatePrefixes(g, prefixes)
	return prefixes
}
