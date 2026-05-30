//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what buildEndpoint / buildRequest / buildResponse / ensureRequest / convertFieldTypes / applyJsonProperty / extractOneField / extractClassFields / extractEnumValues / fieldsToFormParams / extractResponseStatus / extractResponseStatusArg / extractReturnInfo / assignReturnTypeInfo / extractOneRoute / extractMethodEndpoints 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func qFileInfo(t *testing.T, javaSrc string) *fileInfo {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{
		absPath:     "/abs/R.java",
		relPath:     "R.java",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		imports:     map[string]string{},
	}
}

func TestBuildEndpoint(t *testing.T) {
	ri := resourceInfo{prefix: "/users", roles: []string{"admin"}}
	ep := endpointInfo{method: "GET", path: "/{id}", handler: "get", params: []scanner.Param{{Name: "id", Type: "string"}}, returnType: "UserDto"}
	got := buildEndpoint(ri, ep)
	if got.Path != "/users/{id}" || got.Method != "GET" {
		t.Fatalf("meta: %+v", got)
	}
	if len(got.Roles) != 1 || got.Roles[0] != "admin" {
		t.Fatalf("roles: %v", got.Roles)
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
	r := buildRequest(endpointInfo{bodyType: "UserDto", bodyVarName: "dto"})
	if r == nil || r.Body == nil || r.Body.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}

func TestBuildRequest_FormFields(t *testing.T) {
	r := buildRequest(endpointInfo{formParams: []scanner.Param{{Name: "name"}}})
	if r == nil || len(r.FormFields) != 1 {
		t.Fatalf("got %+v", r)
	}
}

func TestBuildResponse_Nil(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestBuildResponse_DefaultStatus(t *testing.T) {
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}

func TestBuildResponse_ExplicitStatusArray(t *testing.T) {
	r := buildResponse(endpointInfo{method: "GET", statusCode: "200", returnType: "UserDto", returnIsArray: true})
	if r == nil || r.Status != "200" || r.Body != "array" {
		t.Fatalf("got %+v", r)
	}
}

func TestEnsureRequest(t *testing.T) {
	ep := &scanner.Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	prev := ep.Request
	ensureRequest(ep)
	if ep.Request != prev {
		t.Fatal("should not replace existing request")
	}
}

func TestConvertFieldTypes(t *testing.T) {
	fields := []scanner.Field{{Name: "id", Type: "Long"}, {Name: "name", Type: "String"}}
	got := convertFieldTypes(fields)
	if got[0].Type != "integer" || got[1].Type != "string" {
		t.Fatalf("got %+v", got)
	}
}

func TestApplyJsonProperty(t *testing.T) {
	root, _ := parseJava([]byte(`class D { @JsonProperty("user_name") private String userName; }`))
	src := []byte(`class D { @JsonProperty("user_name") private String userName; }`)
	field := findAllByType(root, "field_declaration")[0]
	f := &scanner.Field{Name: "userName"}
	applyJsonProperty(field, src, f)
	if f.JSON != "user_name" {
		t.Fatalf("got %q", f.JSON)
	}
}

func TestApplyJsonProperty_None(t *testing.T) {
	root, _ := parseJava([]byte(`class D { private String name; }`))
	src := []byte(`class D { private String name; }`)
	field := findAllByType(root, "field_declaration")[0]
	f := &scanner.Field{}
	applyJsonProperty(field, src, f)
	if f.JSON != "" {
		t.Fatalf("expected empty, got %q", f.JSON)
	}
}

func TestExtractOneFieldAndClassFields(t *testing.T) {
	fi := qFileInfo(t, `class UserDto {
		@JsonProperty("user_name") private String name;
		@NotNull private int age;
		private static String CONST;
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	fields := extractClassFields(cls, fi.src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields (static skipped), got %d: %+v", len(fields), fields)
	}
	if fields[0].Name != "name" || fields[0].JSON != "user_name" {
		t.Fatalf("field0: %+v", fields[0])
	}
}

func TestExtractEnumValues(t *testing.T) {
	fi := qFileInfo(t, `enum Status { OPEN, CLOSED, PENDING }`)
	cls := findAllByType(fi.root, "enum_declaration")[0]
	vals := extractEnumValues(cls, fi.src)
	if len(vals) != 3 || vals[0] != "OPEN" {
		t.Fatalf("got %v", vals)
	}
}

func TestFieldsToFormParams(t *testing.T) {
	fields := []scanner.Field{
		{Name: "name", Type: "string"},
		{Name: "userName", JSON: "user_name", Type: "string"},
	}
	params := fieldsToFormParams(fields)
	if len(params) != 2 || params[1].Name != "user_name" {
		t.Fatalf("got %+v", params)
	}
}

func TestExtractResponseStatus(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@POST public Response create() { return Response.status(201).build(); }
	}`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractResponseStatus(m, fi.src, ep)
	if ep.statusCode != "201" {
		t.Fatalf("got %q", ep.statusCode)
	}
}

func TestExtractReturnInfo(t *testing.T) {
	fi := qFileInfo(t, `class R { public List<UserDto> list() { return null; } }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractReturnInfo(m, fi.src, ep)
	if ep.returnType != "UserDto" || !ep.returnIsArray {
		t.Fatalf("got %+v", ep)
	}
}

func TestAssignReturnTypeInfo_UnknownType(t *testing.T) {
	ep := endpointInfo{returnType: "SomeDto"}
	resp := &scanner.Response{}
	assignReturnTypeInfo(ep, resp)
	// SomeDto -> openAPIType object, so TypeName set
	if resp.TypeName != "SomeDto" {
		t.Fatalf("got %+v", resp)
	}
}

func TestExtractOneRoute(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@GET @Path("/{id}") public UserDto get(@PathParam("id") Long id) { return null; }
	}`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok {
		t.Fatal("expected route")
	}
	if ep.method != "GET" || ep.path != "/{id}" || ep.handler != "get" {
		t.Fatalf("got %+v", ep)
	}
	if len(ep.params) != 1 {
		t.Fatalf("params: %+v", ep.params)
	}
}

func TestExtractOneRoute_NotRoute(t *testing.T) {
	fi := qFileInfo(t, `class R { public void helper() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	if _, ok := extractOneRoute(m, fi); ok {
		t.Fatal("expected false for non-route method")
	}
}

func TestExtractMethodEndpoints(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@GET public String a() { return ""; }
		@POST public String b() { return ""; }
		public void helper() {}
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}
