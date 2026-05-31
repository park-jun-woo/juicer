//ff:func feature=scan type=test topic=nestjs control=iteration dimension=1
//ff:what isBuiltinTSType 빌트인/원시 타입 판별 테스트
package nestjs

import "testing"

func TestIsBuiltinTSType(t *testing.T) {
	for _, s := range []string{"string", "number", "boolean", "Date", "any", "Array", "Promise", "Record"} {
		if !isBuiltinTSType(s) {
			t.Errorf("%q should be builtin", s)
		}
	}
	for _, s := range []string{"UserDto", "Album", ""} {
		if isBuiltinTSType(s) {
			t.Errorf("%q should not be builtin", s)
		}
	}
}
