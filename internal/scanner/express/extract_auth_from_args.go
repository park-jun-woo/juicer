//ff:func feature=scan type=extract control=sequence topic=express
//ff:what argNodes를 순회하여 AuthLevel과 Roles를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractAuthFromArgs(argNodes []*sitter.Node, src []byte) (string, []string) {
	if len(argNodes) < 2 {
		return "public", nil
	}
	return extractAuthFromMiddlewareNodes(argNodes[1:len(argNodes)-1], src)
}
