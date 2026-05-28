//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 변수명으로 같은 파일 내 배열 리터럴 노드를 찾아 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func findArrayLiteral(root *sitter.Node, src []byte, varName string) *sitter.Node {
	for _, vd := range findAllByType(root, "variable_declarator") {
		if !declaratorMatchesName(vd, src, varName) {
			continue
		}
		arr := findChildByType(vd, "array")
		if arr != nil {
			return arr
		}
	}
	return nil
}
