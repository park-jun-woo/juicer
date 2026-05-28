//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what express() 및 express.Router() 인스턴스 변수를 수집한다
package express

func collectRouters(fi *fileInfo, expressRouterAliases map[string]bool) map[string]bool {
	routers := make(map[string]bool)
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		collectRouterFromDecl(decl, fi, routers, expressRouterAliases)
	}
	paramRouters := collectParamRouters(fi)
	for name := range paramRouters {
		routers[name] = true
	}
	return routers
}
