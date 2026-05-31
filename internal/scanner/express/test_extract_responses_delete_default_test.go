//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractResponses_DeleteDefault: res 호출 없는 DELETE 핸들러는 204 기본값을 받는다 (Phase140)
package express

import "testing"

func TestExtractResponses_DeleteDefault(t *testing.T) {
	src := []byte(`const h = (req, res) => { doStuff(); };`)
	fi := mustParse(t, src)
	ri := routeInfo{Method: "DELETE", HandlerNode: firstArrow(t, fi)}
	resps := extractResponses(fi, ri)
	if len(resps) != 1 || resps[0].Status != "204" {
		t.Fatalf("got %+v, want 204", resps)
	}
}
