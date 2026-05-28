//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 클래스 노드에서 controllerInfo를 구성한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func buildControllerInfo(cls *sitter.Node, fi *fileInfo) controllerInfo {
	ci := controllerInfo{
		file:    fi.relPath,
		absFile: fi.absPath,
		imports: fi.imports,
	}
	nameNode := findChildByType(cls, "identifier")
	if nameNode != nil {
		ci.className = nodeText(nameNode, fi.src)
	}
	ci.prefix = extractClassPrefix(cls, fi.src)
	ci.roles = extractClassRoles(cls, fi.src)
	ci.interfaces = extractInterfaces(cls, fi.src)
	ci.endpoints = extractMethodEndpoints(cls, fi)
	if len(ci.endpoints) == 0 && len(ci.interfaces) > 0 {
		resolveControllerInterfaceEndpoints(&ci, fi)
	}
	return ci
}
