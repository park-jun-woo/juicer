//ff:func feature=scan type=test control=sequence topic=django
//ff:what decorators 테스트 헬퍼
package django

import sitter "github.com/smacker/go-tree-sitter"

func decorators(root *sitter.Node) []*sitter.Node {
	return findAllByType(root, "decorator")
}
