//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 메서드 노드에서 HTTP 라우트를 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractOneRoute(m *sitter.Node, fi *fileInfo) (endpointInfo, bool) {
	method, path, found := extractHTTPMethodAndPath(m, fi.src)
	if !found {
		return endpointInfo{}, false
	}
	nameNode := findChildByType(m, "identifier")
	handler := ""
	if nameNode != nil {
		handler = nodeText(nameNode, fi.src)
	}
	ep := endpointInfo{
		method:  method,
		path:    path,
		handler: handler,
		file:    fi.relPath,
		line:    int(m.StartPoint().Row) + 1,
	}
	ep.roles = extractAuthorizeRoles(m, fi.src)
	extractMethodParams(m, fi.src, &ep)
	extractReturnInfo(m, fi.src, &ep)
	return ep, true
}
