//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what 가드(!== POST && !== GET) + POST 블록에서 GET에 body 미전파, POST에 body events 확인
package supafunc

import (
	"testing"
)

func TestStateGuardPattern(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/state/index.ts", `
Deno.serve(async (req) => {
  if (req.method !== 'POST' && req.method !== 'GET') {
    return new Response('Method not allowed', { status: 405 })
  }
  const url = new URL(req.url)
  const id = url.searchParams.get("session_id")
  if (req.method === 'POST') {
    const { events } = await req.json()
    await handleEvents(events)
  }
  return new Response(JSON.stringify({ state: "ok" }), { status: 200 })
})
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	getEp := findEndpoint(result.Endpoints, "GET", "/state")
	if getEp == nil {
		t.Fatal("missing GET /state")
	}
	if getEp.Request != nil && getEp.Request.Body != nil {
		t.Fatalf("GET /state should have no body, got fields: %+v", getEp.Request.Body.Fields)
	}

	postEp := findEndpoint(result.Endpoints, "POST", "/state")
	if postEp == nil {
		t.Fatal("missing POST /state")
	}
	if postEp.Request == nil || postEp.Request.Body == nil {
		t.Fatal("POST /state should have body")
	}
	if len(postEp.Request.Body.Fields) != 1 || postEp.Request.Body.Fields[0].Name != "events" {
		t.Fatalf("POST /state body fields: %+v", postEp.Request.Body.Fields)
	}
}
