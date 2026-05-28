//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 어노테이션에서 key에 해당하는 int 값을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func annotationIntValue(ann *sitter.Node, src []byte, key string) (int, bool) {
	args := annotationArgs(ann, src)
	if args == nil {
		return 0, false
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "element_value_pair" {
			continue
		}
		if v, ok := extractElementPairIntValue(child, src, key); ok {
			return v, true
		}
	}
	return 0, false
}
