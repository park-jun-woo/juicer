//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what 멀티-메서드 order 함수에서 POST → items, PATCH → id,status 분리 확인
package supafunc

import (
	"testing"
)

func TestMethodBodySeparation_Order(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/order/index.ts", `
Deno.serve(async (req) => {
  if (req.method === "POST") {
    const { items, upsells, discounts } = await req.json()
    return new Response(JSON.stringify({ created: true }), { status: 201 })
  }
  if (req.method === "PATCH") {
    const { id, status } = await req.json()
    return new Response(JSON.stringify({ updated: true }), { status: 200 })
  }
})
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	postEp := findEndpoint(result.Endpoints, "POST", "/order")
	if postEp == nil {
		t.Fatal("missing POST /order")
	}
	if postEp.Request == nil || postEp.Request.Body == nil {
		t.Fatal("POST /order should have body")
	}
	postFields := map[string]bool{}
	for _, f := range postEp.Request.Body.Fields {
		postFields[f.Name] = true
	}
	if !postFields["items"] || !postFields["upsells"] || !postFields["discounts"] {
		t.Fatalf("POST /order should have items, upsells, discounts; got %+v", postEp.Request.Body.Fields)
	}
	if postFields["id"] || postFields["status"] {
		t.Fatalf("POST /order should NOT have PATCH fields; got %+v", postEp.Request.Body.Fields)
	}

	patchEp := findEndpoint(result.Endpoints, "PATCH", "/order")
	if patchEp == nil {
		t.Fatal("missing PATCH /order")
	}
	if patchEp.Request == nil || patchEp.Request.Body == nil {
		t.Fatal("PATCH /order should have body")
	}
	patchFields := map[string]bool{}
	for _, f := range patchEp.Request.Body.Fields {
		patchFields[f.Name] = true
	}
	if !patchFields["id"] || !patchFields["status"] {
		t.Fatalf("PATCH /order should have id, status; got %+v", patchEp.Request.Body.Fields)
	}
	if patchFields["items"] || patchFields["upsells"] || patchFields["discounts"] {
		t.Fatalf("PATCH /order should NOT have POST fields; got %+v", patchEp.Request.Body.Fields)
	}
}
