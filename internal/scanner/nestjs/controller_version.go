//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 클래스 노드에서 @Controller 데코레이터의 version 프로퍼티를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// controllerVersion returns the version from @Controller({ version: '1' }).
func controllerVersion(cls *sitter.Node, src []byte) string {
	decorators := findDecorators(cls, src)
	for _, d := range decorators {
		if d.name == DecController && d.objectProps != nil {
			return d.objectProps["version"]
		}
	}
	return ""
}
