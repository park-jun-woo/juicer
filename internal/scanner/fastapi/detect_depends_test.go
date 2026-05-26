//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestExtractDependsFuncName 테스트
package fastapi

import "testing"

func TestExtractDependsFuncName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Depends(get_current_user)", "get_current_user"},
		{"Depends(auth)", "auth"},
		{"Depends()", ""},
		{"NotDepends(foo)", ""},
	}
	for _, tc := range tests {
		got := extractDependsFuncName(tc.input)
		if got != tc.want {
			t.Errorf("extractDependsFuncName(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
