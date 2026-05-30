//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what isAnnotatedDepends: alias / inline Annotated+Depends / 비해당
package fastapi

import "testing"

func TestIsAnnotatedDepends_Alias(t *testing.T) {
	if !isAnnotatedDepends("SessionDep", map[string]string{"SessionDep": "get_db"}) {
		t.Fatal("expected true for alias")
	}
}

func TestIsAnnotatedDepends_Inline(t *testing.T) {
	if !isAnnotatedDepends("Annotated[User, Depends(get_user)]", nil) {
		t.Fatal("expected true for inline")
	}
}

func TestIsAnnotatedDepends_False(t *testing.T) {
	if isAnnotatedDepends("int", nil) {
		t.Fatal("expected false")
	}
	if isAnnotatedDepends("Annotated[int, Query()]", map[string]string{"x": "y"}) {
		t.Fatal("expected false for Annotated without Depends")
	}
}
