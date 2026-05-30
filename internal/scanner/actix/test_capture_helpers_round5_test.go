//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestCaptureHelpers_Round5 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCaptureHelpers_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::scope("/api"); }`)
	call := aFirst(t, fi.root, "call_expression")
	var rootName string
	captureCallRoot(call, fi.src, &rootName)

	var arg string
	walkNodes(fi.root, func(n *sitter.Node) {
		captureScopedCallArg(n, fi.src, "scope", &arg)
	})
}
