//ff:func feature=scan type=test control=iteration dimension=1 topic=scanner
//ff:what isPrimitiveTypeName 테스트 (round5)
package scanner

import "testing"

func TestIsPrimitiveTypeName_Round5(t *testing.T) {
	for _, p := range []string{"bool", "int", "float", "str", "boolean", "number", "string", "any", "dict"} {
		if !isPrimitiveTypeName(p) {
			t.Errorf("%s should be primitive", p)
		}
	}
	for _, np := range []string{"User", "MyDTO", "", "List"} {
		if isPrimitiveTypeName(np) {
			t.Errorf("%s should not be primitive", np)
		}
	}
}
