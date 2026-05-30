//ff:func feature=scan type=test control=sequence topic=spring
//ff:what round5 미커버 함수 직접 호출 테스트 (spring)
package spring

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func sFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
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

func sParse(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func sParam(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := sParse(t, src)
	return sFirst(t, root, "formal_parameter"), b
}

func TestAnnotationArgs_Round5(t *testing.T) {
	root, src := sParse(t, `@RequestMapping("/x") class C {}`)
	ann := sFirst(t, root, "annotation")
	if annotationArgs(ann, src) == nil {
		t.Fatal("expected args")
	}
}

func TestResolveAnnotationName_Round5(t *testing.T) {
	root, src := sParse(t, `@PathVariable("id") class C {}`)
	ann := sFirst(t, root, "annotation")
	if got := resolveAnnotationName(ann, src, "fb"); got != "id" {
		t.Fatalf("got %q", got)
	}
	if got := resolveAnnotationName(nil, src, "fb"); got != "fb" {
		t.Fatalf("nil: %q", got)
	}
}

func TestExtractElementPair_Round5(t *testing.T) {
	root, src := sParse(t, `@Foo(name = "bar", code = 404) class C {}`)
	pairs := findAllByType(root, "element_value_pair")
	var sval string
	var ival int
	for _, p := range pairs {
		if v, ok := extractElementPairValue(p, src, "name"); ok {
			sval = v
		}
		if v, ok := extractElementPairIntValue(p, src, "code"); ok {
			ival = v
		}
	}
	if sval != "bar" {
		t.Errorf("name: %q", sval)
	}
	if ival != 404 {
		t.Errorf("code: %d", ival)
	}
}

func TestExtractParamNameAndType_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(String userName) {} }`)
	if got := extractParamName(param, src); got != "userName" {
		t.Errorf("name: %q", got)
	}
	if got := extractParamType(param, src); got != "String" {
		t.Errorf("type: %q", got)
	}
}

func TestExtractRoleStrings_Round5(t *testing.T) {
	root, src := sParse(t, `@Secured({"ROLE_ADMIN", "ROLE_USER"}) class C {}`)
	ann := sFirst(t, root, "annotation")
	args := annotationArgs(ann, src)
	roles := extractRoleStrings(args, src)
	if len(roles) != 2 {
		t.Fatalf("roles: %v", roles)
	}
}

func TestExtractClassRoles_Round5(t *testing.T) {
	root, src := sParse(t, `@Secured("ROLE_ADMIN") class C {}`)
	cls := sFirst(t, root, "class_declaration")
	roles := extractClassRoles(cls, src)
	if len(roles) != 1 {
		t.Fatalf("roles: %v", roles)
	}
}

func TestExtractStatusArgAndFromArgs_Round5(t *testing.T) {
	root, src := sParse(t, `class C { void m() { ResponseEntity.status(HttpStatus.CREATED).build(); } }`)
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
	got := extractStatusArg(inv, src)
	if got == "" {
		t.Fatalf("status arg empty")
	}

	// direct extractStatusFromArgs on a @ResponseStatus(HttpStatus.CREATED) annotation
	root2, src2 := sParse(t, `@ResponseStatus(HttpStatus.CREATED) class C {}`)
	ann := sFirst(t, root2, "annotation")
	args := annotationArgs(ann, src2)
	if args == nil {
		t.Fatal("no annotation args")
	}
	if s := extractStatusFromArgs(args, src2); s == "" {
		t.Fatalf("extractStatusFromArgs empty")
	}
}

func TestExtractBodyStatus_Round5(t *testing.T) {
	// method body returning ResponseEntity.status(HttpStatus.CREATED)
	root, src := sParse(t, `class C {
		public Object create() {
			return ResponseEntity.status(HttpStatus.CREATED).build();
		}
	}`)
	m := sFirst(t, root, "method_declaration")
	if got := extractBodyStatus(m, src); got == "" {
		t.Fatalf("expected a status code, got empty")
	}
	// method with no body status -> empty
	root2, src2 := sParse(t, `class C { public Object g() { return null; } }`)
	m2 := sFirst(t, root2, "method_declaration")
	if got := extractBodyStatus(m2, src2); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestClassifyPathVariable_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@PathVariable("id") Long id) {} }`)
	var ep endpointInfo
	classifyPathVariable(param, src, &ep, "Long", "id")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("params: %+v", ep.params)
	}
}

func TestClassifyRequestHeader_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestHeader("X-Token") String token) {} }`)
	var ep endpointInfo
	classifyRequestHeader(param, src, &ep, "String", "token")
	if len(ep.headers) != 1 {
		t.Fatalf("headers: %+v", ep.headers)
	}
}

func TestClassifyRequestParam_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestParam("q") String q) {} }`)
	var ep endpointInfo
	classifyRequestParam(param, src, &ep, "String", "q", map[string]string{}, "C.java", "/abs")
	if len(ep.query) != 1 || ep.query[0].Name != "q" {
		t.Fatalf("query: %+v", ep.query)
	}
}

func TestClassifyRequestPart_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestPart("file") MultipartFile file) {} }`)
	var ep endpointInfo
	classifyRequestPart(param, src, &ep, "file")
	if len(ep.files) != 1 {
		t.Fatalf("files: %+v", ep.files)
	}
}

func TestExtractReturnInfo_And_AssignReturnTypeInfo_Round5(t *testing.T) {
	root, src := sParse(t, `class C { public java.util.List<UserDto> all() { return null; } }`)
	m := sFirst(t, root, "method_declaration")
	var ep endpointInfo
	extractReturnInfo(m, src, &ep)
	if ep.returnType == "" {
		t.Fatalf("returnType empty: %+v", ep)
	}
	var resp scanner.Response
	assignReturnTypeInfo(ep, &resp)
	if resp.TypeName == "" {
		t.Fatalf("resp typename empty: %+v", resp)
	}
}

func TestBuildControllerEndpoints_Round5(t *testing.T) {
	ci := controllerInfo{
		prefix:    "/users",
		className: "UserController",
		file:      "C.java",
		endpoints: []endpointInfo{
			{method: "GET", path: "/{id}", handler: "get", returnType: "UserDto"},
		},
		imports: map[string]string{},
	}
	eps, _ := buildControllerEndpoints(ci, "/abs", 0)
	if len(eps) != 1 || eps[0].Path != "/users/{id}" {
		t.Fatalf("endpoints: %+v", eps)
	}
}

func TestResolveClassFromAST_Round5(t *testing.T) {
	root, src := sParse(t, `class UserDto { public String name; public int age; }`)
	fields := resolveClassFromAST(root, src, "UserDto", "C.java", "/abs", map[string][]scanner.Field{})
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d: %+v", len(fields), fields)
	}
}

func TestResolveClassFromASTWithParams_Round5(t *testing.T) {
	root, src := sParse(t, `class Page<T> { public T content; public int total; }`)
	fields, params := resolveClassFromASTWithParams(root, src, "Page", "C.java", "/abs", map[string][]scanner.Field{})
	if len(fields) == 0 {
		t.Fatalf("fields empty")
	}
	_ = params
}

func TestResolveClassFields_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "UserDto.java"), []byte(`package x;
class UserDto { public String name; }
`), 0o644); err != nil {
		t.Fatal(err)
	}
	fields, err := resolveClassFields(filepath.Join(dir, "UserDto.java"), "UserDto", dir, map[string][]scanner.Field{})
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("fields: %+v", fields)
	}
}

func TestResolveSameFileInterface_Round5(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "Api.java")
	if err := os.WriteFile(p, []byte(`interface UserApi {}
class UserController implements UserApi {}
`), 0o644); err != nil {
		t.Fatal(err)
	}
	got := resolveSameFileInterface(p, "UserApi")
	if got != p {
		t.Fatalf("expected same file %q, got %q", p, got)
	}
	if got := resolveSameFileInterface(p, "Missing"); got != "" {
		t.Fatalf("expected empty for missing iface, got %q", got)
	}
}

func TestFindInterfaceFile_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "UserApi.java"), []byte(`package com.x;
interface UserApi {}
`), 0o644); err != nil {
		t.Fatal(err)
	}
	imports := map[string]string{"UserApi": "com.x.UserApi"}
	got := findInterfaceFile("UserApi", imports, filepath.Join(dir, "UserController.java"), dir)
	if got == "" {
		t.Fatalf("expected to find interface file")
	}
}

func TestExtractInterfaceMethodEndpoints_Round5(t *testing.T) {
	fi := sFileInfo(t, `
@RequestMapping("/api")
interface UserApi {
    @GetMapping("/users")
    java.util.List<UserDto> list();
}
`)
	iface := sFirst(t, fi.root, "interface_declaration")
	eps := extractInterfaceMethodEndpoints(iface, fi)
	if len(eps) == 0 {
		t.Fatalf("expected interface endpoints, got %d", len(eps))
	}
}

func TestResolveControllerInterfaceEndpoints_Round5(t *testing.T) {
	fi := sFileInfo(t, `
interface UserApi {
    @GetMapping("/u")
    UserDto get();
}
@RestController
class UserController implements UserApi {
    public UserDto get() { return null; }
}
`)
	ci := &controllerInfo{
		className:  "UserController",
		file:       fi.relPath,
		absFile:    fi.absPath,
		interfaces: []string{"UserApi"},
		imports:    map[string]string{},
	}
	resolveControllerInterfaceEndpoints(ci, fi)
	// should not panic; endpoints may be populated from the interface
}
