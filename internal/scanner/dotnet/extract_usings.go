//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what using 지시문에서 네임스페이스를 수집한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractUsings(root *sitter.Node, src []byte) []string {
	var result []string
	for i := 0; i < int(root.ChildCount()); i++ {
		child := root.Child(i)
		if child.Type() != "using_directive" {
			continue
		}
		ns := extractUsingNamespace(child, src)
		if ns != "" {
			result = append(result, ns)
		}
	}
	return result
}
