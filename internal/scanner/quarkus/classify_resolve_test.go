//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what classifyParam 계열 / single_int_arg / annotation_args / element pair / unquote / generic / substitute / unwrap / roles / walk / resolve 테스트
package quarkus

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstParam(t *testing.T, javaSrc string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no formal_parameter")
	}
	return params[0], b
}

func TestClassifyParam_Path(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@PathParam("id") String id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}

func TestClassifyParam_Query(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@QueryParam("limit") int limit) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.query) != 1 || ep.query[0].Name != "limit" {
		t.Fatalf("got %+v", ep.query)
	}
}

func TestClassifyParam_Header(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@HeaderParam("X-Token") String tok) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.headers) != 1 {
		t.Fatalf("got %+v", ep.headers)
	}
}

func TestClassifyParam_Form(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@FormParam("name") String name) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.formParams) != 1 {
		t.Fatalf("got %+v", ep.formParams)
	}
}

func TestClassifyParam_RestForm(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@RestForm("name") String name) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.formParams) != 1 {
		t.Fatalf("got %+v", ep.formParams)
	}
}

func TestClassifyParam_BeanParam(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@BeanParam Filters f) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.formType != "Filters" {
		t.Fatalf("formType: %q", ep.formType)
	}
}

func TestClassifyParam_Body(t *testing.T) {
	p, src := firstParam(t, `class R { void m(UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.bodyType != "UserDto" {
		t.Fatalf("bodyType: %q", ep.bodyType)
	}
}

func TestClassifyQueryParam_Default(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@DefaultValue("5") @QueryParam("limit") int limit) {} }`)
	ep := &endpointInfo{}
	classifyQueryParam(p, src, ep, "int", "limit")
	if len(ep.query) != 1 || ep.query[0].Default != "5" {
		t.Fatalf("got %+v", ep.query)
	}
}

func TestSingleIntArg(t *testing.T) {
	root, _ := parseJava([]byte(`class C { @Min(5) int x; }`))
	src := []byte(`class C { @Min(5) int x; }`)
	ann := findAllByType(root, "annotation")[0]
	v, ok := singleIntArg(ann, src)
	if !ok || v != 5 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestUnquoteJava(t *testing.T) {
	if unquoteJava(`"hi"`) != "hi" {
		t.Fatal("quoted")
	}
	if unquoteJava("x") != "x" {
		t.Fatal("short")
	}
	if unquoteJava("noq") != "noq" {
		t.Fatal("unquoted")
	}
}

func TestStripGeneric(t *testing.T) {
	if stripGeneric("List<String>") != "List" {
		t.Fatal("generic")
	}
	if stripGeneric("String") != "String" {
		t.Fatal("plain")
	}
}

func TestSplitGenericArgs(t *testing.T) {
	got := splitGenericArgs("String, Map<String,Integer>, Long")
	if len(got) != 3 || got[1] != "Map<String,Integer>" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractGenericArgs(t *testing.T) {
	if got := extractGenericArgs("Map<String, Long>"); got != "String, Long" {
		t.Fatalf("got %q", got)
	}
	if got := extractGenericArgs("String"); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestSubstituteType(t *testing.T) {
	m := map[string]string{"T": "string"}
	if got := substituteType("T", m); got != "string" {
		t.Fatalf("plain: %q", got)
	}
	if got := substituteType("[]T", m); got != "[]string" {
		t.Fatalf("slice: %q", got)
	}
	if got := substituteType("array:T", m); got != "array:string" {
		t.Fatalf("array: %q", got)
	}
	if got := substituteType("List<T>", m); got != "List<string>" {
		t.Fatalf("generic: %q", got)
	}
	if got := substituteType("Other", m); got != "Other" {
		t.Fatalf("unmapped: %q", got)
	}
}

func TestSubstituteTypeParams(t *testing.T) {
	fields := []scanner.Field{{Name: "data", Type: "T"}}
	got := substituteTypeParams(fields, []string{"T"}, []string{"string"})
	if got[0].Type != "string" {
		t.Fatalf("got %+v", got)
	}
	// no type args -> unchanged
	same := substituteTypeParams(fields, nil, nil)
	if same[0].Type != "T" {
		t.Fatalf("unchanged: %+v", same)
	}
}

func TestUnwrapReturnType(t *testing.T) {
	if _, ok := unwrapReturnType("void"); ok {
		t.Fatal("void")
	}
	if _, ok := unwrapReturnType("Response"); ok {
		t.Fatal("Response")
	}
	// Uni<UserDto> unwraps to single UserDto (ok=false means non-array)
	if got, ok := unwrapReturnType("Uni<UserDto>"); ok || got != "UserDto" {
		t.Fatalf("Uni: %q %v", got, ok)
	}
	// List<UserDto> -> array (ok=true)
	if got, ok := unwrapReturnType("List<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("List: %q %v", got, ok)
	}
	// Multi<UserDto> -> array (ok=true)
	if got, ok := unwrapReturnType("Multi<UserDto>"); !ok || got != "UserDto" {
		t.Fatalf("Multi: %q %v", got, ok)
	}
}

func TestExtractRolesAllowed(t *testing.T) {
	root, _ := parseJava([]byte(`@RolesAllowed({"admin"}) class R {}`))
	src := []byte(`@RolesAllowed({"admin"}) class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractRolesAllowed(cls, src)
	if len(roles) != 1 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractClassRoles_Authenticated(t *testing.T) {
	root, _ := parseJava([]byte(`@Authenticated class R {}`))
	src := []byte(`@Authenticated class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractClassRoles(cls, src)
	if len(roles) != 1 || roles[0] != "**" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractClassRoles_None(t *testing.T) {
	root, _ := parseJava([]byte(`class R {}`))
	src := []byte(`class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if roles := extractClassRoles(cls, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}

func TestWalkNodes(t *testing.T) {
	root, _ := parseJava([]byte(`class C { void a() {} void b() {} }`))
	count := 0
	walkNodes(root, func(n *sitter.Node) {
		if n.Type() == "method_declaration" {
			count++
		}
	})
	if count != 2 {
		t.Fatalf("got %d", count)
	}
}

func TestExtractIntLiteralFromArgList(t *testing.T) {
	root, _ := parseJava([]byte(`class C { void m() { status(201); } }`))
	src := []byte(`class C { void m() { status(201); } }`)
	argLists := findAllByType(root, "argument_list")
	if len(argLists) == 0 {
		t.Skip("no arg list")
	}
	for _, al := range argLists {
		if got := extractIntLiteralFromArgList(al, src); got == "201" {
			return
		}
	}
	t.Fatal("did not find 201")
}

func TestFindModifiers(t *testing.T) {
	root, _ := parseJava([]byte(`class C { public int x; }`))
	fields := findAllByType(root, "field_declaration")
	if findModifiers(fields[0]) == nil {
		t.Fatal("expected modifiers")
	}
}
