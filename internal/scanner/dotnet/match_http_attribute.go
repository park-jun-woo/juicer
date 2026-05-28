//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트 리스트에서 HTTP 메서드 어트리뷰트를 매칭한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func matchHTTPAttribute(attrList *sitter.Node, src []byte) (string, string, bool) {
	for _, attr := range childrenOfType(attrList, "attribute") {
		name := attributeName(attr, src)
		httpMethod, ok := httpMethodAttributes[name]
		if !ok {
			continue
		}
		path := attributeFirstStringArg(attr, src)
		return httpMethod, path, true
	}
	return "", "", false
}
