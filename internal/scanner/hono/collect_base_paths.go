//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what new Hono().basePath("/api") 호출에서 변수명→basePath 매핑을 수집한다
package hono

func collectBasePaths(fi *fileInfo, honoVars map[string]bool) map[string]string {
	basePaths := make(map[string]string)
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		collectBasePathFromDecl(decl, fi, basePaths)
	}
	return basePaths
}
