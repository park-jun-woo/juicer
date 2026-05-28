//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 메서드 노드에서 HTTP 라우트를 추출한다
package spring

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
	ep.roles = extractRolesFromNode(m, fi.src)
	extractMethodParams(m, fi.src, &ep, fi.imports, fi.absPath, fi.projectRoot)
	extractReturnInfo(m, fi.src, &ep)
	extractResponseStatus(m, fi.src, &ep)
	if ep.statusCode == "" {
		ep.statusCode = extractBodyStatus(m, fi.src)
	}
	return ep, true
}
