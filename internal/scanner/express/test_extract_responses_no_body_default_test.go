//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractResponses_NoBodyDefault 테스트
package express

import "testing"

func TestExtractResponses_NoBodyDefault(t *testing.T) {
	fi := mustParse(t, []byte(`const x=1;`))
	resps := extractResponses(fi, routeInfo{Handler: "(anonymous)"})
	if len(resps) != 1 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}
