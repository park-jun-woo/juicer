//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what extractDependsFuncName 테스트
package fastapi

import "testing"

func TestExtractDependsFuncName_All(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Depends(get_current_user)", "get_current_user"},
		{"Depends( foo )", "foo"},
		{"Query(default=5)", ""},
		{"", ""},
	}
	for _, tt := range tests {
		got := extractDependsFuncName(tt.in)
		if got != tt.want {
			t.Errorf("extractDependsFuncName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
