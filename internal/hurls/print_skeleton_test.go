//ff:func feature=hurl type=render control=sequence
//ff:what TestPrintSkeleton_Basic 테스트
package hurls

import (
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestPrintSkeleton_Basic(t *testing.T) {
	ep := &scanner.Endpoint{
		Method: "GET",
		Path:   "/api/health",
	}
	printSkeleton(ep, "/tmp/tests")
}
