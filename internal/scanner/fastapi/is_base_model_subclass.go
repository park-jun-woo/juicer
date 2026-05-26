//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스가 BaseModel을 상속하는지 확인한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// isBaseModelSubclass checks if a class inherits from BaseModel.
func isBaseModelSubclass(cls *sitter.Node, src []byte) bool {
	args := findChildByType(cls, "argument_list")
	if args == nil {
		return false
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		text := nodeText(child, src)
		if text == "BaseModel" || strings.HasSuffix(text, ".BaseModel") {
			return true
		}
	}
	return false
}
