//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what 단일 메서드 함수에서 기존 동작 유지 회귀 테스트
package supafunc

import (
	"testing"
)

func TestMethodBodySeparation_SingleMethod(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "supabase/functions/hello/index.ts", `
Deno.serve(async (req) => {
  const { name, age } = await req.json()
  return new Response(JSON.stringify({ message: "Hello " + name }), {
    headers: { "Content-Type": "application/json" },
    status: 200,
  })
})
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Method != "POST" {
		t.Fatalf("expected POST, got %s", ep.Method)
	}
	if ep.Path != "/hello" {
		t.Fatalf("expected /hello, got %s", ep.Path)
	}
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("should have request body")
	}
	fields := map[string]bool{}
	for _, f := range ep.Request.Body.Fields {
		fields[f.Name] = true
	}
	if !fields["name"] || !fields["age"] {
		t.Fatalf("expected name and age fields, got %+v", ep.Request.Body.Fields)
	}
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "200" {
		t.Fatalf("expected response status 200, got %+v", ep.Responses)
	}
}
