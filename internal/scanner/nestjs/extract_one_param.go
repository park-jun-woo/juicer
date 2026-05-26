//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 단일 파라미터의 데코레이터를 해석하여 파라미터 종류별로 분류한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractOneParam processes a single formal parameter with decorators.
func extractOneParam(param *sitter.Node, src []byte, result *methodParams) {
	decorators := paramDecorators(param, src)
	paramName, paramType := paramNameAndType(param, src)
	applyParamDecorators(decorators, paramName, paramType, result)
}
