//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_Concat 테스트
package fiber

import "testing"

func TestCollectStringParts_Concat(t *testing.T) {

	got := collectFor(t, `"/api" + "/v1" + baseURL + opts.X`)
	if len(got) != 2 || got[0] != "/api" || got[1] != "/v1" {
		t.Fatalf("got %v, want [/api /v1]", got)
	}
}
