//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractResponses_MultipleDedup 테스트
package express

import "testing"

func TestExtractResponses_MultipleDedup(t *testing.T) {
	src := []byte(`const h = (req, res) => { res.json({}); res.json({}); res.status(404).send('x'); };`)
	fi := mustParse(t, src)
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	resps := extractResponses(fi, ri)

	if len(resps) != 2 {
		t.Fatalf("expected 2 distinct responses, got %+v", resps)
	}
}
