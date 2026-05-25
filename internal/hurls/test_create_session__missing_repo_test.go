//ff:func feature=hurl type=session control=sequence
//ff:what TestCreateSession_MissingRepo 테스트
package hurls

import "testing"

func TestCreateSession_MissingRepo(t *testing.T) {
	err := createSession("http://localhost", "tests", "")
	if err == nil {
		t.Fatal("expected error")
	}
}
