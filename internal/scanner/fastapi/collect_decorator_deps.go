//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터 목록에서 dependencies=[Depends(func)] 미들웨어를 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectDecoratorDeps iterates decorators and collects Depends function names
// from dependencies= keyword arguments.
func collectDecoratorDeps(decorators []*sitter.Node, src []byte) []string {
	var deps []string
	for _, dec := range decorators {
		callNode, _ := findDecoratorNodes(dec)
		found := extractDecoratorDeps(callNode, src)
		deps = append(deps, found...)
	}
	return deps
}
