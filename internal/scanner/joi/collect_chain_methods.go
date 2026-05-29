//ff:func feature=scan type=parse control=iteration dimension=1 topic=joi
//ff:what Joi 메서드 체인에서 모든 메서드를 inner→outer 순서로 수집한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

// CollectChainMethods — 메서드 체인을 inner→outer 순서로 수집
func CollectChainMethods(node *sitter.Node, src []byte) []ChainMethod {
	var methods []ChainMethod
	collectChainMethodsRecursive(node, src, &methods)
	for i, j := 0, len(methods)-1; i < j; i, j = i+1, j-1 {
		methods[i], methods[j] = methods[j], methods[i]
	}
	return methods
}
