//ff:func feature=scan type=extract control=sequence topic=express
//ff:what app.use()의 인자에서 prefix와 router 변수명을 파싱한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func parseUseMountArgs(args *sitter.Node, src []byte, imports map[string]string) *useMount {
	argNodes := collectArgNodes(args)
	if len(argNodes) < 2 {
		return nil
	}
	prefixNode := argNodes[0]
	if prefixNode.Type() != "string" {
		return nil
	}
	prefix := unquoteTS(nodeText(prefixNode, src))
	routerNode := argNodes[len(argNodes)-1]
	if routerNode.Type() != "identifier" {
		return nil
	}
	varName := nodeText(routerNode, src)
	filePath := imports[varName]
	return &useMount{
		Prefix:   prefix,
		VarName:  varName,
		FilePath: filePath,
	}
}
