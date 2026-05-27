//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스 정의에서 부모 클래스 이름 목록을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectParentNames extracts parent class names from a class definition's argument list.
func collectParentNames(cls *sitter.Node, src []byte) []string {
	args := findChildByType(cls, "argument_list")
	if args == nil {
		return nil
	}
	var names []string
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		// keyword_argument (e.g., table=True)와 구문 토큰(,())은 부모 이름이 아니다
		if child.Type() == "keyword_argument" {
			continue
		}
		text := nodeText(child, src)
		if text != "" && text != "," && text != "(" && text != ")" {
			names = append(names, text)
		}
	}
	return names
}
