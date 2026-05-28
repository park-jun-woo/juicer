//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what new Hono() 인스턴스 변수명을 수집한다
package hono

func collectHonoVars(fi *fileInfo) map[string]bool {
	vars := make(map[string]bool)
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		collectHonoVarFromDecl(decl, fi, vars)
	}
	return vars
}
