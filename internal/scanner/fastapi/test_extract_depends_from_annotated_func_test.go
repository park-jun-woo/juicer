//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractDependsFromAnnotated_Func 테스트
package fastapi

import "testing"

func TestExtractDependsFromAnnotated_Func(t *testing.T) {
	if got := extractDependsFromAnnotated("Annotated[User, Depends(get_user)]"); got != "get_user" {
		t.Fatalf("got %q", got)
	}
}
