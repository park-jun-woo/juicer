//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 어노테이션에서 인자 목록 노드를 반환한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func annotationArgs(ann *sitter.Node, src []byte) *sitter.Node {
	return findChildByType(ann, "annotation_argument_list")
}
