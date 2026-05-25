//ff:func feature=hurl type=session control=sequence
//ff:what TestCreateSession_MissingTests 테스트
package hurls

import "testing"

func TestCreateSession_MissingTests(t *testing.T) {
	err := createSession("http://localhost", "", "repo")
	if err == nil {
		t.Fatal("expected error")
	}
}
