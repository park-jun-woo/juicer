//ff:func feature=scan type=extract control=selection topic=django
//ff:what path() 호출의 두 번째 인자를 view 이름 또는 include로 해석한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// resolveSecondArg resolves the second argument of a path() call.
func resolveSecondArg(entry *urlEntry, arg *sitter.Node, src []byte) {
	switch arg.Type() {
	case "call":
		resolveCallArg(entry, arg, src)
	case "identifier":
		entry.viewName = nodeText(arg, src)
	case "attribute":
		entry.viewName = nodeText(arg, src)
	}
}
