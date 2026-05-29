//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what require('@fastify/autoload') 형태의 변수명을 names 집합에 추가한다
package fastify

func collectAutoloadRequireNames(fi *fileInfo, names map[string]bool) {
	for _, declarator := range findAllByType(fi.Root, "variable_declarator") {
		if v := autoloadRequireName(declarator, fi.Src); v != "" {
			names[v] = true
		}
	}
}
