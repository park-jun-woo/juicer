//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what round5 미커버 함수 직접 호출 테스트 (quarkus)
package quarkus

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func qFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == typ {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}

func qParse(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func TestAnnotationName_Round5(t *testing.T) {
	root, src := qParse(t, "@GET class C {}")
	ann := qFirst(t, root, "marker_annotation")
	if got := annotationName(ann, src); got != "GET" {
		t.Fatalf("got %q", got)
	}
}

func TestAnnotationArgs_Round5(t *testing.T) {
	root, src := qParse(t, `@Path("/x") class C {}`)
	ann := qFirst(t, root, "annotation")
	args := annotationArgs(ann, src)
	if args == nil {
		t.Fatal("expected args node")
	}
}

func TestResolveAnnotationName_Round5(t *testing.T) {
	root, src := qParse(t, `@PathParam("id") class C {}`)
	ann := qFirst(t, root, "annotation")
	if got := resolveAnnotationName(ann, src, "fallback"); got != "id" {
		t.Fatalf("got %q", got)
	}
	// nil annotation -> fallback
	if got := resolveAnnotationName(nil, src, "fallback"); got != "fallback" {
		t.Fatalf("nil fallback: %q", got)
	}
}

func TestExtractElementPairValue_Round5(t *testing.T) {
	root, src := qParse(t, `@RolesAllowed(value = "admin") class C {}`)
	pair := qFirst(t, root, "element_value_pair")
	v, ok := extractElementPairValue(pair, src, "value")
	if !ok || v != "admin" {
		t.Fatalf("got %q %v", v, ok)
	}
	if _, ok := extractElementPairValue(pair, src, "other"); ok {
		t.Fatal("expected miss for wrong key")
	}
}

func TestExtractElementPairIntValue_Round5(t *testing.T) {
	root, src := qParse(t, `@Max(status = 404) class C {}`)
	pair := qFirst(t, root, "element_value_pair")
	v, ok := extractElementPairIntValue(pair, src, "status")
	if !ok || v != 404 {
		t.Fatalf("got %d %v", v, ok)
	}
}

func TestExtractRolesFromNode_Round5(t *testing.T) {
	root, src := qParse(t, `@RolesAllowed({"admin", "user"}) class C {}`)
	cls := qFirst(t, root, "class_declaration")
	roles := extractRolesFromNode(cls, src)
	if len(roles) != 2 {
		t.Fatalf("expected 2 roles, got %v", roles)
	}
}

func TestExtractOneField_Round5(t *testing.T) {
	root, src := qParse(t, `class Dto { public String name; }`)
	field := qFirst(t, root, "field_declaration")
	f := extractOneField(field, src)
	if f.Name != "name" {
		t.Fatalf("field: %+v", f)
	}
}

func TestExtractResponseStatusArg_Round5(t *testing.T) {
	root, src := qParse(t, `class C { void m() { Response.status(201).build(); } }`)
	// find the status(201) invocation
	var inv *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if inv != nil {
			return
		}
		if n.Type() == "method_invocation" {
			name := n.ChildByFieldName("name")
			if name != nil && nodeText(name, src) == "status" {
				inv = n
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if inv == nil {
		t.Fatal("no status invocation")
	}
	got := extractResponseStatusArg(inv, src)
	if got != "201" {
		t.Fatalf("status: %q", got)
	}
}

func classifyFixture(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := qParse(t, src)
	param := qFirst(t, root, "formal_parameter")
	return param, b
}

func TestClassifyPathParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@PathParam("id") String id) {} }`)
	var ep endpointInfo
	classifyPathParam(param, src, &ep, "String", "id")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("params: %+v", ep.params)
	}
}

func TestClassifyHeaderParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@HeaderParam("X-Token") String token) {} }`)
	var ep endpointInfo
	classifyHeaderParam(param, src, &ep, "String", "token")
	if len(ep.headers) != 1 {
		t.Fatalf("headers: %+v", ep.headers)
	}
}

func TestClassifyFormParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@FormParam("f") String f) {} }`)
	var ep endpointInfo
	classifyFormParam(param, src, &ep, "String", "f")
	if len(ep.formParams) == 0 {
		t.Fatalf("formFields: %+v", ep.formParams)
	}
}

func TestClassifyRestForm_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@RestForm String f) {} }`)
	var ep endpointInfo
	classifyRestForm(param, src, &ep, "String", "f")
	if len(ep.formParams) == 0 {
		t.Fatalf("formFields: %+v", ep.formParams)
	}
}

func TestBuildResourceEndpoints_Round5(t *testing.T) {
	ri := resourceInfo{
		prefix: "/users",
		endpoints: []endpointInfo{
			{method: "GET", path: "/{id}", handler: "get", returnType: "UserDto"},
		},
	}
	eps, _ := buildResourceEndpoints(ri, "/abs", 0)
	if len(eps) != 1 || eps[0].Path != "/users/{id}" {
		t.Fatalf("endpoints: %+v", eps)
	}
}
