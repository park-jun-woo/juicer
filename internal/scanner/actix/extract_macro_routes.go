//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what proc-macro 어트리뷰트(#[get("/path")])에서 라우트를 추출한다
package actix

func extractMacroRoutes(fi *fileInfo) []macroRoute {
	var routes []macroRoute
	var pendingAttrs []macroRoute

	root := fi.root
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		routes, pendingAttrs = consumeMacroChild(child, fi, routes, pendingAttrs)
	}

	return routes
}
