//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsAnnotatedDepends_Alias 테스트
package fastapi

import "testing"

func TestIsAnnotatedDepends_Alias(t *testing.T) {
	if !isAnnotatedDepends("SessionDep", map[string]string{"SessionDep": "get_db"}) {
		t.Fatal("expected true for alias")
	}
}
