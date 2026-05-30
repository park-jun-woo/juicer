//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildRequest_FormFields 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestBuildRequest_FormFields(t *testing.T) {
	r := buildRequest(endpointInfo{formParams: []scanner.Param{{Name: "name"}}})
	if r == nil || len(r.FormFields) != 1 {
		t.Fatalf("got %+v", r)
	}
}
