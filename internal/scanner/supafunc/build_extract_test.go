//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what buildRequest / buildResponses / buildBodyFields / extractEndpoint / extractParamFromCall / extractStatusFromResponse / extractMethodFromCondition / collectMethodsFromCondition 테스트
package supafunc

import (
	"testing"
)

func TestBuildRequest_Nil(t *testing.T) {
	if r := buildRequest(nil, nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestBuildRequest_BodyAndQuery(t *testing.T) {
	r := buildRequest([]string{"task", "status"}, []string{"limit"})
	if r == nil || r.Body == nil || len(r.Body.Fields) != 2 {
		t.Fatalf("body: %+v", r)
	}
	if len(r.Query) != 1 || r.Query[0].Name != "limit" {
		t.Fatalf("query: %+v", r.Query)
	}
}

func TestBuildResponses(t *testing.T) {
	resps := buildResponses([]string{"200", "404"})
	if len(resps) != 2 || resps[0].Status != "200" || resps[0].Kind != "json" {
		t.Fatalf("got %+v", resps)
	}
}

func TestBuildBodyFields(t *testing.T) {
	fields := buildBodyFields([]string{"a", "b"})
	if len(fields) != 2 || fields[0].Name != "a" || fields[0].Type != "unknown" {
		t.Fatalf("got %+v", fields)
	}
}

func TestExtractParamFromCall(t *testing.T) {
	fi := mustParse(t, []byte(`url.searchParams.get("limit");`))
	calls := findAllByType(fi.Root, "call_expression")
	for _, c := range calls {
		if got := extractParamFromCall(c, fi.Src); got == "limit" {
			return
		}
	}
	t.Fatal("did not find limit")
}

func TestExtractStatusFromResponse(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Response(JSON.stringify(x), { status: 404 });`))
	news := findAllByType(fi.Root, "new_expression")
	if len(news) == 0 {
		t.Fatal("no new_expression")
	}
	if got := extractStatusFromResponse(news[0], fi.Src); got != "404" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractStatusFromResponse_None(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Response("ok");`))
	news := findAllByType(fi.Root, "new_expression")
	if got := extractStatusFromResponse(news[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractMethodFromCondition(t *testing.T) {
	fi := mustParse(t, []byte(`if (req.method === "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if len(conds) == 0 {
		t.Fatal("no condition")
	}
	if got := extractMethodFromCondition(conds[0], fi.Src); got != "POST" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractMethodFromCondition_Negated(t *testing.T) {
	// guard pattern (!==) should not register the method
	fi := mustParse(t, []byte(`if (req.method !== "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if got := extractMethodFromCondition(conds[0], fi.Src); got != "" {
		t.Fatalf("expected empty for negated, got %q", got)
	}
}

func TestCollectMethodsFromCondition(t *testing.T) {
	fi := mustParse(t, []byte(`if (req.method === "GET" || req.method === "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if len(conds) == 0 {
		t.Fatal("no condition")
	}
	methods := collectMethodsFromCondition(conds[0], fi.Src)
	if len(methods) != 2 {
		t.Fatalf("got %v", methods)
	}
}

func TestExtractEndpoint_NoCallback(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "hello/index.ts", `const x = 1;`)
	fi, err := parseFile(dir + "/hello/index.ts")
	if err != nil {
		t.Fatal(err)
	}
	eps := extractEndpoint(fi, dir)
	if len(eps) != 1 || eps[0].Method != "POST" || eps[0].Path != "/hello" {
		t.Fatalf("got %+v", eps)
	}
}

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
