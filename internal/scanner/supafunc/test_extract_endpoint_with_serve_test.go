//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestExtractEndpoint_WithServe 테스트
package supafunc

import "testing"

func TestExtractEndpoint_WithServe(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "tasks/index.ts", `
serve(async (req) => {
  if (req.method === "GET") {
    return new Response(JSON.stringify([]))
  }
  return new Response(JSON.stringify({}), { status: 201 })
})
`)
	fi, err := parseFile(dir + "/tasks/index.ts")
	if err != nil {
		t.Fatal(err)
	}
	eps := extractEndpoint(fi, dir)
	if len(eps) == 0 {
		t.Fatalf("expected endpoints, got %+v", eps)
	}
	for _, e := range eps {
		if e.Path != "/tasks" {
			t.Errorf("path %q", e.Path)
		}
	}
}
