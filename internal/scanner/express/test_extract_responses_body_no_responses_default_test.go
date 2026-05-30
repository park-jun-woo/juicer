//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractResponses_BodyNoResponsesDefault 테스트
package express

import "testing"

func TestExtractResponses_BodyNoResponsesDefault(t *testing.T) {
	src := []byte(`const h = (req, res) => { doStuff(); };`)
	fi := mustParse(t, src)
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	resps := extractResponses(fi, ri)
	if len(resps) != 1 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}
