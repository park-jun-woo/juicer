//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 파라미터의 타입 어노테이션이 express.Router 또는 Router인지 확인하여 등록한다
package express

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func checkParamTypeAnnotation(param *sitter.Node, src []byte, routers map[string]bool) {
	if param.Type() != "required_parameter" && param.Type() != "optional_parameter" {
		return
	}
	nameNode := findChildByType(param, "identifier")
	typeAnn := findChildByType(param, "type_annotation")
	if nameNode == nil || typeAnn == nil {
		return
	}
	typeText := nodeText(typeAnn, src)
	typeText = strings.TrimPrefix(typeText, ":")
	typeText = strings.TrimSpace(typeText)
	if typeText == "express.Router" || typeText == "Router" {
		routers[nodeText(nameNode, src)] = true
	}
}
