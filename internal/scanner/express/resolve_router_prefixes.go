//ff:func feature=scan type=extract control=iteration dimension=2 topic=express
//ff:what 마운트 그래프를 풀어 각 라우터 인스턴스(file,var)의 전체 prefix 목록을 계산한다
package express

import "sort"

// routerEdge는 부모 라우터에서 자식 라우터로의 마운트(세그먼트)다.
type routerEdge struct {
	child routerKey
	seg   string
}

// resolveRouterPrefixes는 모든 마운트 엔트리와 라우터 집합으로부터
// 라우터 인스턴스별 전체 prefix 목록(map[routerKey][]string)을 계산한다.
//
// 파일 단위가 아니라 (file, varName) 단위로 키잉하므로:
//   - 같은 라우터를 두 prefix에 마운트하면 두 prefix가 모두 보존된다.
//   - 한 파일에 라우터가 여럿이면 각각 다른 prefix를 받는다.
func resolveRouterPrefixes(mounts []mountEntry, allRouters map[string]map[string]bool) map[routerKey][]string {
	edges := map[routerKey][]routerEdge{}
	indeg := map[routerKey]int{}
	nodes := map[routerKey]bool{}
	for file, rs := range allRouters {
		for v := range rs {
			nodes[routerKey{file, v}] = true
		}
	}

	addEdge := func(parent, child routerKey, seg string) {
		nodes[parent] = true
		nodes[child] = true
		edges[parent] = append(edges[parent], routerEdge{child, seg})
		indeg[child]++
	}

	for _, m := range mounts {
		parent := routerKey{m.sourceFile, m.sourceRouter}
		if m.filePath == "" {
			// 인라인 마운트: 자식은 같은 파일의 로컬 라우터 변수.
			if m.varName == "" {
				continue
			}
			addEdge(parent, routerKey{m.sourceFile, m.varName}, m.prefix)
			continue
		}
		// 크로스파일 마운트.
		if cv := resolveChildVar(m.filePath, m.varName, allRouters); cv != "" {
			addEdge(parent, routerKey{m.filePath, cv}, m.prefix)
			continue
		}
		// 모호: 자식 파일의 모든 라우터에 적용 (파일 단위 폴백).
		childRouters := make([]string, 0, len(allRouters[m.filePath]))
		for v := range allRouters[m.filePath] {
			childRouters = append(childRouters, v)
		}
		sort.Strings(childRouters)
		for _, v := range childRouters {
			addEdge(parent, routerKey{m.filePath, v}, m.prefix)
		}
	}

	// 루트(들어오는 마운트 없음)는 prefix "".
	prefixes := map[routerKey][]string{}
	for n := range nodes {
		if indeg[n] == 0 {
			prefixes[n] = []string{""}
		}
	}

	// 결정적 순서로 부모를 순회하며 수렴할 때까지 전파한다.
	parents := make([]routerKey, 0, len(edges))
	for p := range edges {
		parents = append(parents, p)
	}
	sort.Slice(parents, func(i, j int) bool {
		if parents[i].file != parents[j].file {
			return parents[i].file < parents[j].file
		}
		return parents[i].varName < parents[j].varName
	})

	const maxIter = 100
	for it := 0; it < maxIter; it++ {
		changed := false
		for _, p := range parents {
			pps, ok := prefixes[p]
			if !ok {
				continue
			}
			for _, e := range edges[p] {
				for _, pp := range pps {
					if appendUniquePrefix(prefixes, e.child, joinExpressPath(pp, e.seg)) {
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}
	return prefixes
}

// appendUniquePrefix는 key의 prefix 목록에 v가 없으면 추가하고 true를 반환한다.
func appendUniquePrefix(m map[routerKey][]string, key routerKey, v string) bool {
	for _, existing := range m[key] {
		if existing == v {
			return false
		}
	}
	m[key] = append(m[key], v)
	return true
}
