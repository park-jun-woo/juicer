//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 클래스 노드에서 controllerInfo를 구성한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func buildControllerInfo(cls *sitter.Node, fi *fileInfo) controllerInfo {
	ci := controllerInfo{
		file:    fi.relPath,
		absFile: fi.absPath,
		usings:  fi.usings,
	}
	nameNode := findChildByType(cls, "identifier")
	if nameNode != nil {
		ci.className = nodeText(nameNode, fi.src)
	}
	ci.prefix = extractClassRoute(cls, fi.src, ci.className)
	ci.roles = extractClassRoles(cls, fi.src)
	ci.endpoints = extractMethodEndpoints(cls, fi)
	return ci
}
