//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsAnnotatedDepends_Inline 테스트
package fastapi

import "testing"

func TestIsAnnotatedDepends_Inline(t *testing.T) {
	if !isAnnotatedDepends("Annotated[User, Depends(get_user)]", nil) {
		t.Fatal("expected true for inline")
	}
}
