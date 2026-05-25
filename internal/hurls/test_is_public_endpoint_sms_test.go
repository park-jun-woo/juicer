//ff:func feature=hurl type=parse control=sequence
//ff:what TestIsPublicEndpoint_SMS 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestIsPublicEndpoint_SMS(t *testing.T) {
	ep := scanner.Endpoint{Method: "POST", Path: "/sms"}
	if !isPublicEndpoint(ep) {
		t.Fatal("expected true")
	}
}
