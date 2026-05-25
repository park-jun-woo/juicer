//ff:func feature=hurl type=render control=sequence
//ff:what TestHandleFail_WithStderr 테스트
package hurls

import "testing"

func TestHandleFail_WithStderr(t *testing.T) {
	handleFail("some error output")
}
