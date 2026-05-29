//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 모든 route 그룹에 대해 prefix를 1회 전파하고 변경 여부를 반환한다 (fixpoint 1패스)
package hono

func runPrefixPass(groups []routeGroup, prefixMap map[string]string, honoVars map[string]map[string]bool, imports map[string]map[string]string) bool {
	changed := false
	for _, g := range groups {
		if applyGroupPrefix(g, prefixMap, honoVars, imports) {
			changed = true
		}
	}
	return changed
}
