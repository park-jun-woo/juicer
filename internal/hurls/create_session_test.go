//ff:func feature=hurl type=session control=sequence
//ff:what TestCreateSession_MissingHost 테스트
package hurls

import "testing"

func TestCreateSession_MissingHost(t *testing.T) {
	err := createSession("", "tests", "repo")
	if err == nil {
		t.Fatal("expected error")
	}
}
