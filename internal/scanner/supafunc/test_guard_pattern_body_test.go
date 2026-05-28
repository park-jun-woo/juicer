//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what 가드 패턴(req.method !== "POST")에서 POST만 body를 갖고 다른 메서드는 body가 없는지 확인
package supafunc

import (
	"testing"
)

func TestGuardPatternBody(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/web-checkout/index.ts", `
Deno.serve(async (req) => {
  if (req.method !== 'POST') {
    return new Response('Method not allowed', { status: 405 })
  }
  const { store, items, upsells, client_last_updated } = await req.json()
  return new Response(JSON.stringify({ ok: true }), { status: 200 })
})
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	postEp := findEndpoint(result.Endpoints, "POST", "/web-checkout")
	if postEp == nil {
		t.Fatal("missing POST /web-checkout")
	}
	if postEp.Request == nil || postEp.Request.Body == nil {
		t.Fatal("POST /web-checkout should have body")
	}
	fields := map[string]bool{}
	for _, f := range postEp.Request.Body.Fields {
		fields[f.Name] = true
	}
	if !fields["store"] || !fields["items"] || !fields["upsells"] || !fields["client_last_updated"] {
		t.Fatalf("POST /web-checkout missing expected body fields; got %+v", postEp.Request.Body.Fields)
	}
}
