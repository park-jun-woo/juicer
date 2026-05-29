//ff:func feature=scan type=parse control=sequence topic=joi
//ff:what property 노드와 call_expression에서 ChainMethod를 생성한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

func buildChainMethodFromProp(callNode, prop *sitter.Node, src []byte) ChainMethod {
	cm := ChainMethod{Name: nodeText(prop, src)}
	args := findChildByType(callNode, "arguments")
	if args != nil {
		cm.Args = collectStringArgs(args, src)
	}
	return cm
}
