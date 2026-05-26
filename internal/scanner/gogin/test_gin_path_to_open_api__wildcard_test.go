//ff:func feature=scan type=extract control=sequence
//ff:what TestGinPathToOpenAPI_Wildcard 테스트
package gogin

import "testing"

func TestGinPathToOpenAPI_Wildcard(t *testing.T) {
	got := ginPathToOpenAPI("/files/*path")
	if got != "/files/{path}" {
		t.Fatalf("expected /files/{path}, got %s", got)
	}
}
