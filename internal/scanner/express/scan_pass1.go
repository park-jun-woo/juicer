//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Pass 1: 파일 파싱, 라우터 수집, import 해석, use-prefix 수집
package express

func scanPass1(tsFiles []string, absRoot string) *scanContext {
	parsed := make(map[string]*fileInfo)
	allRouters := make(map[string]map[string]bool)
	aliases := loadTsconfigPaths(absRoot)
	var allMounts []mountEntry
	for _, path := range tsFiles {
		mounts := scanOneFilePass1(path, parsed, allRouters, absRoot, aliases)
		allMounts = append(allMounts, mounts...)
	}
	prefixMap := buildPrefixMap(allMounts)
	return &scanContext{parsed: parsed, allRouters: allRouters, prefixMap: prefixMap, absRoot: absRoot, pathAliases: aliases}
}
