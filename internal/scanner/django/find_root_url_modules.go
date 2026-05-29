//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 다른 모듈의 include 대상이 아닌 루트 urlconf 모듈을 찾는다
package django

// findRootURLModules returns modules that are not included by any other module.
// These are the entry points for URL expansion (root urlconf candidates).
func findRootURLModules(byModule map[string][]urlEntry) []string {
	included := make(map[string]bool)
	for _, entries := range byModule {
		collectIncludedModules(entries, byModule, included)
	}
	var roots []string
	for mod := range byModule {
		if !included[mod] {
			roots = append(roots, mod)
		}
	}
	return roots
}
