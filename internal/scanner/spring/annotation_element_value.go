//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 어노테이션에서 key에 해당하는 element value를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func annotationElementValue(ann *sitter.Node, src []byte, key string) string {
	args := annotationArgs(ann, src)
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "element_value_pair" {
			continue
		}
		if v, ok := extractElementPairValue(child, src, key); ok {
			return v
		}
	}
	return ""
}
