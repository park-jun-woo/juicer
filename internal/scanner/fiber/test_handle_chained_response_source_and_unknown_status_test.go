//ff:func feature=scan type=test control=sequence
//ff:what TestHandleChainedResponse_SourceAndUnknownStatus 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestHandleChainedResponse_SourceAndUnknownStatus(t *testing.T) {
	ep := scanner.Endpoint{}

	sc, oc := chainCall("JSON", nil)
	handleChainedResponse(&ep, sc, oc, "JSON", nil, "respond")
	if len(ep.Responses) != 1 {
		t.Fatalf("expected 1 response")
	}
	if ep.Responses[0].Status != "(unknown)" {
		t.Errorf("status = %q, want (unknown)", ep.Responses[0].Status)
	}
	if ep.Responses[0].Source != "respond" {
		t.Errorf("source = %q, want respond", ep.Responses[0].Source)
	}
}
