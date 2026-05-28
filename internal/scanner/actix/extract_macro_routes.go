//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what proc-macro 어트리뷰트(#[get("/path")])에서 라우트를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type macroRoute struct {
	method   string
	path     string
	handler  string
	funcNode *sitter.Node
	file     *fileInfo
}

func extractMacroRoutes(fi *fileInfo) []macroRoute {
	var routes []macroRoute
	root := fi.root
	src := fi.src

	var pendingAttrs []macroRoute

	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)

		if child.Type() == "attribute_item" {
			r := parseMacroAttribute(child, src)
			if r != nil {
				pendingAttrs = append(pendingAttrs, *r)
			}
			continue
		}

		if child.Type() == "function_item" && len(pendingAttrs) > 0 {
			nameNode := findChildByType(child, "identifier")
			handler := ""
			if nameNode != nil {
				handler = nodeText(nameNode, src)
			}
			for _, attr := range pendingAttrs {
				attr.handler = handler
				attr.funcNode = child
				attr.file = fi
				routes = append(routes, attr)
			}
			pendingAttrs = nil
			continue
		}

		pendingAttrs = nil
	}

	return routes
}

func parseMacroAttribute(attrItem *sitter.Node, src []byte) *macroRoute {
	attr := findChildByType(attrItem, "attribute")
	if attr == nil {
		return nil
	}
	nameNode := findChildByType(attr, "identifier")
	if nameNode == nil {
		return nil
	}
	name := nodeText(nameNode, src)
	method, ok := httpMacros[name]
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
	path := nodeText(strContent, src)

	return &macroRoute{
		method: method,
		path:   path,
	}
}
