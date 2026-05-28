//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what lexical_declaration 내 variable_declarator에서 express/Router 인스턴스를 찾아 등록한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectRouterFromDecl(decl *sitter.Node, fi *fileInfo, routers map[string]bool, aliases map[string]bool) {
	for _, declarator := range findAllByType(decl, "variable_declarator") {
		nameNode := findChildByType(declarator, "identifier")
		if nameNode == nil {
			continue
		}
		varName := nodeText(nameNode, fi.Src)
		valueNode := findInitValue(declarator)
		if valueNode == nil {
			continue
		}
		if isExpressCall(valueNode, fi.Src) || isExpressRouterCall(valueNode, fi.Src) || isRouterStandaloneCall(valueNode, fi.Src, aliases) {
			routers[varName] = true
		}
	}
}
