//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what property 노드와 call_expression에서 ChainMethod를 생성한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func buildChainMethodFromProp(callNode, prop *sitter.Node, src []byte) ChainMethod {
	args := findChildByType(callNode, "arguments")
	cm := ChainMethod{Name: nodeText(prop, src), Node: args}
	if args != nil {
		cm.Args = collectStringArgs(args, src)
	}
	return cm
}
