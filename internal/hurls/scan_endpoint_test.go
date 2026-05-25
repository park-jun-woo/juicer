//ff:func feature=hurl type=parse control=sequence
//ff:what TestScanEndpoint_InvalidDir 테스트
package hurls

import "testing"

func TestScanEndpoint_InvalidDir(t *testing.T) {
	ep := scanEndpoint("/nonexistent/dir", "GET /health")
	if ep != nil {
		t.Fatal("expected nil for invalid dir")
	}
}
