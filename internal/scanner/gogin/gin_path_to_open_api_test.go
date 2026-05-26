//ff:func feature=scan type=test control=sequence
//ff:what TestGinPathToOpenAPI_Param 테스트
package gogin

import "testing"

func TestGinPathToOpenAPI_Param(t *testing.T) {
	got := ginPathToOpenAPI("/api/users/:id")
	if got != "/api/users/{id}" {
		t.Fatalf("expected /api/users/{id}, got %s", got)
	}
}

