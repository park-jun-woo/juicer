//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 클래스가 [ApiController]인지 확인한다
package dotnet

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func isApiController(cls *sitter.Node, src []byte) bool {
	if hasAttribute(cls, src, AttrApiController) {
		return true
	}
	base := findChildByType(cls, "base_list")
	if base == nil {
		return false
	}
	text := nodeText(base, src)
	return strings.Contains(text, "ControllerBase") || strings.Contains(text, "Controller")
}
