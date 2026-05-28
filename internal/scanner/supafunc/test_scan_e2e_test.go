//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestScanE2E 전체 스캔 통합 테스트 (tempdir로 Supabase 프로젝트 구조 모사)
package supafunc

import (
	"testing"
)

func TestScanE2E(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/hello/index.ts", `
Deno.serve(async (req) => {
  const { name } = await req.json()
  return new Response(JSON.stringify({ message: "Hello " + name }), {
    headers: { "Content-Type": "application/json" },
    status: 200,
  })
})
`)

	writeFile(t, dir, "supabase/functions/get-user/index.ts", `
import { serve } from "https://deno.land/std/http/server.ts"

serve(async (req) => {
  if (req.method === "OPTIONS") {
    return new Response("ok", { status: 204 })
  }
  if (req.method === "GET") {
    const url = new URL(req.url)
    const id = url.searchParams.get("id")
    return new Response(JSON.stringify({ user: id }), { status: 200 })
  }
  if (req.method === "POST") {
    const { id, email } = await req.json()
    return new Response(JSON.stringify({ created: true }), { status: 201 })
  }
})
`)

	writeFile(t, dir, "supabase/functions/_shared/cors.ts", `
export const corsHeaders = { "Access-Control-Allow-Origin": "*" }
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	methodPaths := map[string]bool{}
	for _, ep := range result.Endpoints {
		methodPaths[ep.Method+" "+ep.Path] = true
	}

	if !methodPaths["POST /hello"] {
		t.Fatal("missing POST /hello")
	}
	if !methodPaths["GET /get-user"] {
		t.Fatal("missing GET /get-user")
	}
	if !methodPaths["POST /get-user"] {
		t.Fatal("missing POST /get-user")
	}
	if methodPaths["OPTIONS /get-user"] {
		t.Fatal("OPTIONS should be excluded")
	}

	hello := findEndpoint(result.Endpoints, "POST", "/hello")
	if hello == nil || hello.Request == nil || hello.Request.Body == nil {
		t.Fatal("hello should have request body")
	}
	if len(hello.Request.Body.Fields) != 1 || hello.Request.Body.Fields[0].Name != "name" {
		t.Fatalf("hello body fields: %+v", hello.Request.Body.Fields)
	}

	getUser := findEndpoint(result.Endpoints, "GET", "/get-user")
	if getUser == nil || getUser.Request == nil || len(getUser.Request.Query) == 0 {
		t.Fatal("get-user GET should have query params")
	}
	if getUser.Request.Query[0].Name != "id" {
		t.Fatalf("expected query param 'id', got %q", getUser.Request.Query[0].Name)
	}
	if getUser.Request.Body != nil {
		t.Fatalf("get-user GET should have no body, got %+v", getUser.Request.Body.Fields)
	}

	postUser := findEndpoint(result.Endpoints, "POST", "/get-user")
	if postUser == nil || postUser.Request == nil || postUser.Request.Body == nil {
		t.Fatal("get-user POST should have request body")
	}
	postUserFields := map[string]bool{}
	for _, f := range postUser.Request.Body.Fields {
		postUserFields[f.Name] = true
	}
	if !postUserFields["id"] || !postUserFields["email"] {
		t.Fatalf("get-user POST body fields: %+v", postUser.Request.Body.Fields)
	}
}
