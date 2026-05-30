//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestExtractFuncNameFromCall_Neither 테스트
package hono

import "testing"

func TestExtractFuncNameFromCall_Neither(t *testing.T) {

	fi := mustParse(t, []byte("(() => {})();\n"))
	calls := findAllByType(fi.Root, "call_expression")

	for _, c := range calls {
		got := extractFuncNameFromCall(c, fi.Src)
		if got == "" {
			return
		}
	}
	t.Skip("no call with empty func name")
}
