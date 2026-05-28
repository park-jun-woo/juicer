//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트의 정수 인자를 반환한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func attributeIntArgs(attr *sitter.Node, src []byte) []int {
	args := findChildByType(attr, "attribute_argument_list")
	if args == nil {
		return nil
	}
	var result []int
	for _, arg := range childrenOfType(args, "attribute_argument") {
		v, ok := findIntLiteralInArg(arg, src)
		if ok {
			result = append(result, v)
		}
	}
	return result
}
