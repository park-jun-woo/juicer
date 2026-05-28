//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what OR 조건(method === "POST" || method === "PUT")에서 두 메서드 모두 body를 갖고 GET/DELETE는 body 없음 확인
package supafunc

import (
	"testing"
)

func TestOrConditionMultiMethod(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/restful-tasks/index.ts", `
Deno.serve(async (req) => {
  const { url, method } = req
  const taskUrl = new URL(url)
  const id = taskUrl.searchParams.get("id")

  if (method === 'POST' || method === 'PUT') {
    const body = await req.json()
    var task = body.task
  }

  switch (true) {
    case method === 'GET':
      return new Response(JSON.stringify({ tasks: [] }), { status: 200 })
    case method === 'DELETE':
      return new Response(JSON.stringify({ deleted: true }), { status: 200 })
    case method === 'POST':
      return new Response(JSON.stringify({ created: true }), { status: 201 })
    case method === 'PUT':
      return new Response(JSON.stringify({ updated: true }), { status: 200 })
  }
})
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	postEp := findEndpoint(result.Endpoints, "POST", "/restful-tasks")
	if postEp == nil {
		t.Fatal("missing POST /restful-tasks")
	}
	if postEp.Request == nil || postEp.Request.Body == nil {
		t.Fatal("POST /restful-tasks should have body")
	}

	putEp := findEndpoint(result.Endpoints, "PUT", "/restful-tasks")
	if putEp == nil {
		t.Fatal("missing PUT /restful-tasks")
	}
	if putEp.Request == nil || putEp.Request.Body == nil {
		t.Fatal("PUT /restful-tasks should have body")
	}

	getEp := findEndpoint(result.Endpoints, "GET", "/restful-tasks")
	if getEp == nil {
		t.Fatal("missing GET /restful-tasks")
	}
	if getEp.Request != nil && getEp.Request.Body != nil {
		t.Fatalf("GET /restful-tasks should have no body, got fields: %+v", getEp.Request.Body.Fields)
	}

	deleteEp := findEndpoint(result.Endpoints, "DELETE", "/restful-tasks")
	if deleteEp == nil {
		t.Fatal("missing DELETE /restful-tasks")
	}
	if deleteEp.Request != nil && deleteEp.Request.Body != nil {
		t.Fatalf("DELETE /restful-tasks should have no body, got fields: %+v", deleteEp.Request.Body.Fields)
	}
}
