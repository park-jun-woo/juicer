//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 메서드 어트리뷰트에서 HTTP 메서드와 경로를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractHTTPMethodAndPath(m *sitter.Node, src []byte) (string, string, bool) {
	for _, attrList := range childrenOfType(m, "attribute_list") {
		method, path, ok := matchHTTPAttribute(attrList, src)
		if ok {
			return method, path, true
		}
	}
	return "", "", false
}
