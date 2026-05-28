//ff:func feature=scan type=extract control=sequence topic=django
//ff:what call 타입 인자를 include() 또는 .as_view()로 해석한다
package django

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveCallArg resolves a call-type second argument (include() or .as_view()).
func resolveCallArg(entry *urlEntry, arg *sitter.Node, src []byte) {
	callFunc := findChildByType(arg, "identifier")
	callAttr := findChildByType(arg, "attribute")

	if callFunc != nil && nodeText(callFunc, src) == "include" {
		entry.isInclude = true
		innerArgs := findChildByType(arg, "argument_list")
		if innerArgs != nil {
			entry.includeModule = firstStringArg(innerArgs, src)
		}
		return
	}

	if callAttr != nil {
		attrText := nodeText(callAttr, src)
		if strings.HasSuffix(attrText, ".as_view") {
			entry.viewName = strings.TrimSuffix(attrText, ".as_view")
			return
		}
		entry.viewName = attrText
		return
	}

	if callFunc != nil {
		entry.viewName = nodeText(callFunc, src)
	}
}
