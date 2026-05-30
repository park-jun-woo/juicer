//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromAnnotated_Empty 테스트
package fastapi

import "testing"

func TestExtractDependsFromAnnotated_Empty(t *testing.T) {

	got := extractDependsFromAnnotated("Annotated[OAuth2PasswordRequestForm, Depends()]")
	if got != "OAuth2PasswordRequestForm" {
		t.Fatalf("got %q", got)
	}
}
