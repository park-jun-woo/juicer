//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 클래스 노드에서 UseGuards/Roles 데코레이터를 수집하여 controllerInfo에 저장한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectClassLevelDecorators extracts @UseGuards and @Roles from a class
// node's decorators and populates ci.classMiddleware / ci.classRoles.
func collectClassLevelDecorators(cls *sitter.Node, src []byte, ci *controllerInfo) {
	for _, d := range findDecorators(cls, src) {
		if d.name == DecUseGuards {
			ci.classMiddleware = append(ci.classMiddleware, d.args...)
		}
		if d.name == DecRoles {
			ci.classRoles = append(ci.classRoles, d.args...)
		}
	}
}
