//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 그래프의 부모 라우터들을 (file, varName) 기준 결정적 순서로 정렬해 반환한다
package express

import "sort"

func sortedParents(g *mountGraph) []routerKey {
	parents := make([]routerKey, 0, len(g.edges))
	for p := range g.edges {
		parents = append(parents, p)
	}
	sort.Slice(parents, func(i, j int) bool {
		if parents[i].file != parents[j].file {
			return parents[i].file < parents[j].file
		}
		return parents[i].varName < parents[j].varName
	})
	return parents
}
