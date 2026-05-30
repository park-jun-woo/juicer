//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractResponses_None 테스트
package hono

import "testing"

func TestExtractResponses_None(t *testing.T) {
	fi := mustParse(t, []byte(`const h = () => { foo(); };`))
	if resps := extractResponses(fi, 1); len(resps) != 0 {
		t.Fatalf("expected no responses, got %+v", resps)
	}
}
