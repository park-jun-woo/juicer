//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what variable_declarator가 require('@fastify/autoload')이면 변수명을 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func autoloadRequireName(declarator *sitter.Node, src []byte) string {
	nameNode := findChildByType(declarator, "identifier")
	if nameNode == nil {
		return ""
	}
	callNode := findInitCallExpr(declarator)
	if callNode == nil {
		return ""
	}
	fnNode := findChildByType(callNode, "identifier")
	if fnNode == nil || nodeText(fnNode, src) != "require" {
		return ""
	}
	if extractCallStringArg(callNode, src) != autoloadModule {
		return ""
	}
	return nodeText(nameNode, src)
}
