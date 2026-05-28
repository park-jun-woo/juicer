//ff:func feature=scan type=extract control=sequence topic=express
//ff:what function_declaration이 name과 일치하면 statement_block을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func matchFunctionDecl(node *sitter.Node, src []byte, name string) *sitter.Node {
	if node.Type() != "function_declaration" {
		return nil
	}
	nameNode := node.ChildByFieldName("name")
	if nameNode == nil || nodeText(nameNode, src) != name {
		return nil
	}
	return findChildByType(node, "statement_block")
}
