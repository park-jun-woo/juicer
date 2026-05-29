//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 한 모듈의 entry들에서 include로 참조되는 대상 모듈들을 included 집합에 모은다
package django

// collectIncludedModules records, into included, every module referenced by include() in entries.
func collectIncludedModules(entries []urlEntry, byModule map[string][]urlEntry, included map[string]bool) {
	for _, e := range entries {
		if !e.isInclude {
			continue
		}
		if mod, ok := resolveIncludeModule(e.includeModule, byModule); ok {
			included[mod] = true
		}
	}
}
