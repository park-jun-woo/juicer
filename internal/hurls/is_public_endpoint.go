//ff:func feature=hurl type=parse control=sequence
//ff:what 엔드포인트가 인증 불필요한 공개 엔드포인트인지 판별
package hurls

import (
	"strings"

	"github.com/park-jun-woo/juicer/scanner"
)

// isPublicEndpoint returns true if the endpoint requires no authentication.
// Detects health, login, and sms endpoints by path pattern.
func isPublicEndpoint(ep scanner.Endpoint) bool {
	lower := strings.ToLower(ep.Path)
	if strings.HasSuffix(lower, "/health") {
		return true
	}
	if strings.Contains(lower, "/auth/login") {
		return true
	}
	if lower == "/sms" || strings.HasSuffix(lower, "/sms") {
		return true
	}
	return false
}
