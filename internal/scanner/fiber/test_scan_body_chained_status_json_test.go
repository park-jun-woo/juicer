//ff:func feature=scan type=test control=sequence
//ff:what TestScanBody_ChainedStatusJSON 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestScanBody_ChainedStatusJSON(t *testing.T) {
	src := `package m
func h() {
	return c.Status(201).JSON(result)
}
`
	body := handlerBody(t, src)
	ep := &scanner.Endpoint{}
	scanBody(ep, body, "c", nil, &funcIndex{}, "handler")
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "201" {
		t.Fatalf("chained status: %v", ep.Responses)
	}
}
