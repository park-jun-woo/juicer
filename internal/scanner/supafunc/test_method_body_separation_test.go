//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what 멀티-메서드 state 함수에서 GET은 body 없음, POST는 body events 확인
package supafunc

import (
	"testing"
)

func TestMethodBodySeparation_State(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/state/index.ts", `
Deno.serve(async (req) => {
  if (req.method === "GET") {
    const url = new URL(req.url)
    const id = url.searchParams.get("session_id")
    return new Response(JSON.stringify({ state: "ok" }), { status: 200 })
  }
  if (req.method === "POST") {
    const { events } = await req.json()
    return new Response(JSON.stringify({ saved: true }), { status: 201 })
  }
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
	if getEp.Request == nil || len(getEp.Request.Query) == 0 {
		t.Fatal("GET /state should have query params")
	}
	if getEp.Request.Query[0].Name != "session_id" {
		t.Fatalf("expected query param 'session_id', got %q", getEp.Request.Query[0].Name)
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
