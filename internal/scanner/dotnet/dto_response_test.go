//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what DTO 해석 / 응답 추출 함수 테스트
package dotnet

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveDTOType(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } public int Age { get; set; } }`)
	dr := dtoRequest{typeName: "UserDto"}
	fields := resolveDTOType(dr, []*fileInfo{fi}, map[string][]scanner.Field{})
	if len(fields) != 2 || fields[0].Name != "Name" {
		t.Fatalf("got %+v", fields)
	}
}

func TestResolveDTOType_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{"X": {{Name: "a"}}}
	fields := resolveDTOType(dtoRequest{typeName: "X"}, nil, cache)
	if len(fields) != 1 {
		t.Fatalf("got %+v", fields)
	}
}

func TestResolveDTOType_NotFound(t *testing.T) {
	fi := csFileInfo(t, `class Other {}`)
	if fields := resolveDTOType(dtoRequest{typeName: "Missing"}, []*fileInfo{fi}, map[string][]scanner.Field{}); fields != nil {
		t.Fatalf("got %+v", fields)
	}
}

func TestFindClassInFile(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } }`)
	fields := findClassInFile(fi, "UserDto")
	if len(fields) != 1 {
		t.Fatalf("got %+v", fields)
	}
	if findClassInFile(fi, "Missing") != nil {
		t.Fatal("missing should be nil")
	}
}

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ci := controllerInfo{usings: []string{}}
	reqs := collectDTORequests(ep, ci, "/abs", 0)
	if len(reqs) != 2 {
		t.Fatalf("got %d", len(reqs))
	}
}

func TestCollectDTORequests_Primitives(t *testing.T) {
	ep := endpointInfo{bodyType: "string", returnType: "int"}
	if reqs := collectDTORequests(ep, controllerInfo{}, "/abs", 0); len(reqs) != 0 {
		t.Fatalf("got %d", len(reqs))
	}
}

func TestAssignDTOFields(t *testing.T) {
	body := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{}}}
	assignDTOFields(dtoRequest{isBody: true}, body, []scanner.Field{{Name: "x"}})
	if len(body.Request.Body.Fields) != 1 {
		t.Fatalf("body: %+v", body.Request.Body)
	}
	resp := &scanner.Endpoint{Responses: []scanner.Response{{Status: "200"}}}
	assignDTOFields(dtoRequest{}, resp, []scanner.Field{{Name: "y"}})
	if len(resp.Responses[0].Fields) != 1 {
		t.Fatalf("resp: %+v", resp.Responses)
	}
}

func TestAssignReturnTypeInfo(t *testing.T) {
	resp := &scanner.Response{}
	assignReturnTypeInfo(endpointInfo{returnType: "UserDto", returnIsArray: true}, resp)
	if resp.TypeName != "UserDto" || resp.Body != "array" {
		t.Fatalf("got %+v", resp)
	}
}

func TestFieldsToFormParams(t *testing.T) {
	params := fieldsToFormParams([]scanner.Field{{Name: "a", Type: "string"}})
	if len(params) != 1 || params[0].Name != "a" {
		t.Fatalf("got %+v", params)
	}
}

func TestResolveAllDTOs(t *testing.T) {
	fi := csFileInfo(t, `class UserDto { public string Name { get; set; } }`)
	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	dr := dtoRequest{typeName: "UserDto", epIdx: 0}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, []*fileInfo{fi})
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("got %+v", endpoints[0].Responses)
	}
}

func TestIsDIType(t *testing.T) {
	if !isDIType("ILogger<Foo>") {
		t.Fatal("ILogger")
	}
	if !isDIType("AppDbContext") {
		t.Fatal("DbContext")
	}
	if !isDIType("IUserService") {
		t.Fatal("IService")
	}
	if isDIType("UserDto") {
		t.Fatal("dto not DI")
	}
}

func TestExtractParamDefault(t *testing.T) {
	root, src := parseCS(t, `class C { void m(int limit = 10) {} }`)
	params := findAllByType(root, "parameter")
	if got := extractParamDefault(params[0], src); got != "10" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractReturnInfo_Generic(t *testing.T) {
	root, src := parseCS(t, `class C { public ActionResult<UserDto> Get() { return Ok(); } }`)
	m := findAllByType(root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, src, ep)
	if ep.returnType != "UserDto" {
		t.Fatalf("got %q", ep.returnType)
	}
}

func TestExtractReturnInfo_Array(t *testing.T) {
	root, src := parseCS(t, `class C { public ActionResult<List<UserDto>> List() { return Ok(); } }`)
	m := findAllByType(root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, src, ep)
	if ep.returnType != "UserDto" || !ep.returnIsArray {
		t.Fatalf("got %+v", ep)
	}
}
