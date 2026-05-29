//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what #[get("/path")] 같은 HTTP 매크로 어트리뷰트에서 method/path를 파싱한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func parseMacroAttribute(attrItem *sitter.Node, src []byte) *macroRoute {
	attr := findChildByType(attrItem, "attribute")
	if attr == nil {
		return nil
	}
	nameNode := findChildByType(attr, "identifier")
	if nameNode == nil {
		return nil
	}
	method, ok := httpMacros[nodeText(nameNode, src)]
	if !ok {
		return nil
	}

	tokenTree := findChildByType(attr, "token_tree")
	if tokenTree == nil {
		return nil
	}
	strLit := findChildByType(tokenTree, "string_literal")
	if strLit == nil {
		return nil
	}
	strContent := findChildByType(strLit, "string_content")
	if strContent == nil {
		return nil
	}

	return &macroRoute{
		method: method,
		path:   nodeText(strContent, src),
	}
}
