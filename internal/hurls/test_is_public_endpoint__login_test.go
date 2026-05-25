//ff:func feature=hurl type=parse control=sequence
//ff:what TestIsPublicEndpoint_Login 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestIsPublicEndpoint_Login(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/api/auth/login"}
	if !isPublicEndpoint(ep) {
		t.Fatal("expected true")
	}
}
