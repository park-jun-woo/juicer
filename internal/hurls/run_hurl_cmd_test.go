//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunHurlTest_NoHurl 테스트
package hurls

import "testing"

func TestRunHurlTest_NoHurl(t *testing.T) {
	// hurl binary likely not available in test environment
	passed, stderr := runHurlTest("/nonexistent/file.hurl", "http://localhost")
	if passed {
		t.Fatal("expected failure")
	}
	if stderr == "" {
		t.Fatal("expected stderr output")
	}
}
