//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractDependsFromAnnotated: Depends(func) / Depends() / Depends없음 / 미닫힘
package fastapi

import "testing"

func TestExtractDependsFromAnnotated_Func(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[User, Depends(get_user)]"); got != "get_user" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractDependsFromAnnotated_Empty(t *testing.T) {
	// empty Depends() -> falls back to first Annotated type arg
	got := extractDependsFromAnnotated("Annotated[OAuth2PasswordRequestForm, Depends()]")
	if got != "OAuth2PasswordRequestForm" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractDependsFromAnnotated_NoDepends(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[int, Query()]"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractDependsFromAnnotated_NoClose(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[x, Depends("); got != "" {
		t.Fatalf("got %q", got)
	}
}
