//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPathToOpenAPI_Wildcard 테스트
package echo

import "testing"

func TestEchoPathToOpenAPI_Wildcard(t *testing.T) {
	got := echoPathToOpenAPI("/files/*path")
	if got != "/files/{path}" {
		t.Fatalf("expected /files/{path}, got %s", got)
	}
}
