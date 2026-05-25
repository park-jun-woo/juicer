//ff:func feature=hurl type=render control=sequence
//ff:what TestHandleFail_EmptyStderr 테스트
package hurls

import "testing"

func TestHandleFail_EmptyStderr(t *testing.T) {
	handleFail("")
}
