//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what variable_declarator에서 TypeBox Type.Object() 호출의 인자 객체를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractTypeBoxFromDeclarator(declarator *sitter.Node, fi *fileInfo, vars map[string]*sitter.Node) {
	nameNode := findChildByType(declarator, "identifier")
	if nameNode == nil {
		return
	}
	callNode := findInitCallExpr(declarator)
	if callNode == nil || !isTypeBoxObjectCall(callNode, fi.Src) {
		return
	}
	args := findChildByType(callNode, "arguments")
	if args == nil {
		return
	}
	objNode := findChildByType(args, "object")
	if objNode != nil {
		vars[nodeText(nameNode, fi.Src)] = objNode
	}
}
