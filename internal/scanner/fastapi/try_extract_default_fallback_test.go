//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryExtractDefault 단순 케이스 fallback 헬퍼
package fastapi

import "testing"

func tryExtractDefaultFallback(t *testing.T) bool {
	t.Helper()
	src2 := []byte("def f(x=5): pass\n")
	root2, _ := parsePython(src2)
	funcDef2 := findChildByType(root2, "function_definition")
	params2 := findChildByType(funcDef2, "parameters")
	dp := findChildByType(params2, "default_parameter")
	if dp == nil {
		return false
	}
	for i := 0; i < int(dp.ChildCount()); i++ {
		child := dp.Child(i)
		val, _, _ := tryExtractDefault(child, src2)
		if val == "5" {
			return true
		}
	}
	return false
}
