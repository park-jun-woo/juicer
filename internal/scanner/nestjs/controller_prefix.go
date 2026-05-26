//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 클래스 노드에서 @Controller 데코레이터 접두사를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// controllerPrefix returns the path prefix from @Controller('prefix') or ("").
func controllerPrefix(cls *sitter.Node, src []byte) (string, bool) {
	decorators := findDecorators(cls, src)
	for _, d := range decorators {
		if d.name == DecController {
			return d.arg, true
		}
	}
	return "", false
}
