//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestExtractDependsFromAnnotated Annotated 텍스트에서 Depends 함수명 추출 테스트
package fastapi

import "testing"

func TestExtractDependsFromAnnotated(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{"Annotated[Session, Depends(get_db)]", "get_db"},
		{"Annotated[User, Depends(get_current_user)]", "get_current_user"},
		{"Annotated[str, Depends(reusable_oauth2)]", "reusable_oauth2"},
		{"Annotated[OAuth2PasswordRequestForm, Depends()]", "Depends"},
		{"SomeOtherType", ""},
		{"Annotated[str, Header()]", ""},
	}
	for _, tc := range tests {
		got := extractDependsFromAnnotated(tc.text)
		if got != tc.want {
			t.Errorf("extractDependsFromAnnotated(%q): got %q, want %q", tc.text, got, tc.want)
		}
	}
}
