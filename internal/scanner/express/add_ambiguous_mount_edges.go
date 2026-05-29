//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 모호한 크로스파일 마운트를 자식 파일의 모든 라우터에 결정적 순서로 적용한다
package express

import "sort"

func addAmbiguousMountEdges(g *mountGraph, parent routerKey, m mountEntry, allRouters map[string]map[string]bool) {
	childRouters := make([]string, 0, len(allRouters[m.filePath]))
	for v := range allRouters[m.filePath] {
		childRouters = append(childRouters, v)
	}
	sort.Strings(childRouters)
	for _, v := range childRouters {
		graphAddEdge(g, parent, routerKey{m.filePath, v}, m.prefix)
	}
}
