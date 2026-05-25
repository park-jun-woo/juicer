package hurls

import "testing"

func TestHandleFail_WithStderr(t *testing.T) {
	handleFail("some error output")
}

func TestHandleFail_EmptyStderr(t *testing.T) {
	handleFail("")
}
