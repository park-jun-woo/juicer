package hurls

import "testing"

func TestShowTODO_NoScan(t *testing.T) {
	ep := &EndpointStatus{ID: "GET /api/health", Status: "TODO"}
	// repoDir is invalid so scanEndpoint returns nil, falls to basic output
	showTODO(ep, "/nonexistent/repo", "/tmp/tests")
}
