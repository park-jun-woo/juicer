package hurls

import "testing"

func TestScanEndpoint_InvalidDir(t *testing.T) {
	ep := scanEndpoint("/nonexistent/dir", "GET /health")
	if ep != nil {
		t.Fatal("expected nil for invalid dir")
	}
}
