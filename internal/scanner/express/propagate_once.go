//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 모든 부모를 한 번 순회하며 prefix를 전파하고 변경 여부를 반환한다
package express

func propagateOnce(g *mountGraph, parents []routerKey, prefixes map[routerKey][]string) bool {
	changed := false
	for _, p := range parents {
		if propagateParent(g, p, prefixes) {
			changed = true
		}
	}
	return changed
}
