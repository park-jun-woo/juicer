//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestBuildResponses 테스트
package supafunc

import "testing"

func TestBuildResponses(t *testing.T) {
	resps := buildResponses([]string{"200", "404"})
	if len(resps) != 2 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}
