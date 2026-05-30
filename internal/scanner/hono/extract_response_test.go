//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractResponses 테스트
package hono

import "testing"

func TestExtractResponses(t *testing.T) {
	fi := mustParse(t, []byte(`
const h = (c) => {
  foo();
  c.json({ ok: true }, 201);
  c.text("hi");
}
`))
	resps := extractResponses(fi, 1)
	if len(resps) != 2 {
		t.Fatalf("expected 2 responses, got %d: %+v", len(resps), resps)
	}
	if resps[0].Kind != "json" || resps[0].Status != "201" {
		t.Errorf("resp0: %+v", resps[0])
	}
	if resps[1].Kind != "text" {
		t.Errorf("resp1: %+v", resps[1])
	}
}

func TestExtractResponses_None(t *testing.T) {
	fi := mustParse(t, []byte(`const h = () => { foo(); };`))
	if resps := extractResponses(fi, 1); len(resps) != 0 {
		t.Fatalf("expected no responses, got %+v", resps)
	}
}
