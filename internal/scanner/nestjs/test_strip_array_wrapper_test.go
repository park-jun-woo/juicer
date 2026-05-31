//ff:func feature=scan type=test topic=nestjs control=iteration dimension=1
//ff:what stripArrayWrapper T[]/Array<T> 래퍼 제거 테스트
package nestjs

import "testing"

func TestStripArrayWrapper(t *testing.T) {
	cases := []struct {
		in        string
		wantElem  string
		wantArray bool
	}{
		{"UserDto[]", "UserDto", true},
		{"Array<UserDto>", "UserDto", true},
		{"UserDto", "UserDto", false},
		{"string", "string", false},
	}
	for _, c := range cases {
		elem, arr := stripArrayWrapper(c.in)
		if elem != c.wantElem || arr != c.wantArray {
			t.Errorf("stripArrayWrapper(%q) = (%q,%v), want (%q,%v)", c.in, elem, arr, c.wantElem, c.wantArray)
		}
	}
}
