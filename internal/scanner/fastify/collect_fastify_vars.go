//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what lexical_declaration에서 Fastify() 호출로 생성된 인스턴스 변수를 수집한다
package fastify

func collectFastifyVars(fi *fileInfo, instances map[string]bool) {
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		collectFastifyVarFromDecl(decl, fi, instances)
	}
}
