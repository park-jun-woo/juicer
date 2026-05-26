//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 단일 클래스 노드에서 컨트롤러 정보를 생성한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// buildControllerInfo builds a controllerInfo from a class node.
// It returns false if the class does not have a @Controller decorator.
func buildControllerInfo(cls *sitter.Node, src []byte, file string, absFile string, imports map[string]string) (controllerInfo, bool) {
	prefix, ok := controllerPrefix(cls, src)
	if !ok {
		return controllerInfo{}, false
	}
	ci := controllerInfo{
		prefix:  prefix,
		version: controllerVersion(cls, src),
		imports: imports,
	}
	collectClassLevelDecorators(cls, src, &ci)
	direct := extractMethods(cls, src, file)
	inherited := resolveBaseController(cls, src, absFile, imports, file)
	ci.endpoints = mergeEndpoints(inherited, direct)
	applyClassDecorators(&ci)
	return ci, true
}
