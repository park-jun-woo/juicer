//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestIsExternalPackage 테스트
package nestjs

import "testing"

func TestIsExternalPackage(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"./dto/user", false},
		{"../shared/dto", false},
		{"@nestjs/common", true},
		{"@nestjs/mapped-types", true},
		{"@/decorators/field", false},
		{"src/users/dto/create-user.dto", false},
		{"class-validator", true},
		{"rxjs", true},
		{"lib/utils", false},
	}
	for _, tc := range cases {
		got := isExternalPackage(tc.input)
		if got != tc.want {
			t.Errorf("isExternalPackage(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
