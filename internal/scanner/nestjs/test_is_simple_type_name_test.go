//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what isSimpleTypeName 단순 식별자 판별 테스트
package nestjs

import "testing"

func TestIsSimpleTypeName(t *testing.T) {
	for _, s := range []string{"UserDto", "_id", "$ref", "A1"} {
		if !isSimpleTypeName(s) {
			t.Errorf("%q should be simple", s)
		}
	}
	for _, s := range []string{"", "A<B>", "A|B", "A[]", "A.B", "A B"} {
		if isSimpleTypeName(s) {
			t.Errorf("%q should not be simple", s)
		}
	}
}
