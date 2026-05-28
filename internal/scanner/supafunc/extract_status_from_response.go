//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what new Response 노드의 arguments 내 object에서 status 값을 추출한다
package supafunc

import sitter "github.com/smacker/go-tree-sitter"

func extractStatusFromResponse(ne *sitter.Node, src []byte) string {
	args := findChildByType(ne, "arguments")
	if args == nil {
		return ""
	}
	objs := findAllByType(args, "object")
	for _, obj := range objs {
		status := extractStatusFromObject(obj, src)
		if status != "" {
			return status
		}
	}
	return ""
}
