//ff:func feature=scan type=extract control=sequence topic=express
//ff:what variable_declarator의 식별자가 주어진 변수명과 일치하는지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func declaratorMatchesName(vd *sitter.Node, src []byte, varName string) bool {
	id := findChildByType(vd, "identifier")
	return id != nil && nodeText(id, src) == varName
}
