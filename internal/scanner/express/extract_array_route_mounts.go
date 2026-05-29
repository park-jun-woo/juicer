//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what forEach 패턴에서 배열 리터럴의 path/route 쌍을 정적 추출하여 mountEntry 목록을 반환한다
package express

func extractArrayRouteMounts(fi *fileInfo, routers map[string]bool, imports map[string]string, sourcePath string) []mountEntry {
	var entries []mountEntry
	for _, call := range findAllByType(fi.Root, "call_expression") {
		arrVar := matchForEachCall(call, fi.Src, routers)
		if arrVar == "" {
			continue
		}
		arr := findArrayLiteral(fi.Root, fi.Src, arrVar)
		if arr == nil {
			continue
		}
		parentRouter := forEachParentRouter(call, fi.Src, routers)
		for _, pe := range extractObjectEntries(arr, fi.Src) {
			filePath := imports[pe.routeVar]
			entries = append(entries, mountEntry{
				prefix:       pe.path,
				varName:      pe.routeVar,
				filePath:     filePath,
				sourceFile:   sourcePath,
				sourceRouter: parentRouter,
			})
		}
	}
	return entries
}
