//ff:func feature=scan type=test control=sequence topic=spring
//ff:what extractHTTPMethodAndPath / extractClassPrefix / classify* / build* / isRestController / isMultipartFile / springPathToOpenAPI / roles 추출 테스트
package spring

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstMethodS(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	methods := findAllByType(root, "method_declaration")
	if len(methods) == 0 {
		t.Fatal("no method")
	}
	return methods[0], src
}

func firstParamS(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	root, src := parseS(t, javaSrc)
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	return params[0], src
}

func TestExtractHTTPMethodAndPath_GetMapping(t *testing.T) {
	m, src := firstMethodS(t, `class C { @GetMapping("/users") public String list() { return ""; } }`)
	method, path, ok := extractHTTPMethodAndPath(m, src)
	if !ok || method != "GET" || path != "/users" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}

func TestExtractHTTPMethodAndPath_None(t *testing.T) {
	m, src := firstMethodS(t, `class C { public String helper() { return ""; } }`)
	if _, _, ok := extractHTTPMethodAndPath(m, src); ok {
		t.Fatal("expected false")
	}
}

func TestExtractClassPrefix(t *testing.T) {
	root, src := parseS(t, `@RequestMapping("/api") class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPrefix(cls, src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractClassPrefix_None(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPrefix(cls, src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestClassifyParam_PathVariable(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@PathVariable("id") Long id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}

func TestClassifyParam_RequestParam(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestParam("q") String q) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.query) != 1 {
		t.Fatalf("got %+v", ep.query)
	}
}

func TestClassifyParam_RequestBody(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestBody UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.bodyType != "UserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}

func TestClassifyParam_RequestHeader(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestHeader("X-Token") String tok) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.headers) != 1 {
		t.Fatalf("got %+v", ep.headers)
	}
}

func TestClassifyParam_RequestPart(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestPart("file") MultipartFile file) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.files) != 1 || ep.files[0].Type != "string:binary" {
		t.Fatalf("got %+v", ep.files)
	}
}

func TestClassifyParam_ModelAttribute(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@ModelAttribute Filters f) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.formType != "Filters" {
		t.Fatalf("got %q", ep.formType)
	}
}

func TestBuildEndpoint(t *testing.T) {
	ci := controllerInfo{prefix: "/users", roles: []string{"ADMIN"}}
	ep := endpointInfo{method: "GET", path: "/{id}", handler: "get", params: []scanner.Param{{Name: "id", Type: "string"}}, returnType: "UserDto"}
	got := buildEndpoint(ci, ep)
	if got.Path != "/users/{id}" || got.Method != "GET" {
		t.Fatalf("meta: %+v", got)
	}
	if got.Request == nil || len(got.Responses) != 1 {
		t.Fatalf("req/resp: %+v", got)
	}
}

func TestBuildRequest_Nil(t *testing.T) {
	if r := buildRequest(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestBuildRequest_Body(t *testing.T) {
	r := buildRequest(endpointInfo{bodyType: "UserDto"})
	if r == nil || r.Body == nil || r.Body.Method != "RequestBody" {
		t.Fatalf("got %+v", r)
	}
}

func TestBuildResponse(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}

func TestIsRestController(t *testing.T) {
	root, src := parseS(t, `@RestController class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !isRestController(cls, src) {
		t.Fatal("RestController")
	}
	root2, src2 := parseS(t, `@Controller @ResponseBody class D {}`)
	cls2 := findAllByType(root2, "class_declaration")[0]
	if !isRestController(cls2, src2) {
		t.Fatal("Controller+ResponseBody")
	}
	root3, src3 := parseS(t, `class E {}`)
	cls3 := findAllByType(root3, "class_declaration")[0]
	if isRestController(cls3, src3) {
		t.Fatal("plain class")
	}
}

func TestIsMultipartFile(t *testing.T) {
	if !isMultipartFile("MultipartFile") || !isMultipartFile("org.x.MultipartFile") || isMultipartFile("String") {
		t.Fatal("multipart")
	}
}

func TestSpringPathToOpenAPI(t *testing.T) {
	if springPathToOpenAPI("/users/{id}") != "/users/{id}" {
		t.Fatal("path")
	}
}

func TestExtractPreAuthorizeRoles(t *testing.T) {
	root, src := parseS(t, `class C { @PreAuthorize("hasRole('ADMIN')") void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractPreAuthorizeRoles(m, src)
	if len(roles) != 1 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractPreAuthorizeRoles_None(t *testing.T) {
	root, src := parseS(t, `class C { void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	if roles := extractPreAuthorizeRoles(m, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractSecuredRoles(t *testing.T) {
	root, src := parseS(t, `class C { @Secured({"ROLE_ADMIN"}) void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractSecuredRoles(m, src)
	if len(roles) != 1 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractRolesFromNode(t *testing.T) {
	root, src := parseS(t, `class C { @PreAuthorize("hasRole('USER')") void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractRolesFromNode(m, src)
	if len(roles) == 0 || roles[0] != "USER" {
		t.Fatalf("got %v", roles)
	}
}
