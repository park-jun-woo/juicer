//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what collectDTORequests / resolveAllDTOs / resolveDTOType / resolveClassFields(WithParams) / resolveClassFromASTWithParams / resolveEnumFromAST / resolveImportPath / resolveSameFileClass / resolveSamePackageClass / resolveParentFields / mergeParentFields / assignDTOFields 테스트
package quarkus

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ri := resourceInfo{imports: map[string]string{}, absFile: "/abs/R.java"}
	reqs := collectDTORequests(ep, ri, "/abs", 0)
	if len(reqs) != 2 {
		t.Fatalf("expected 2 reqs, got %d", len(reqs))
	}
}

func TestCollectDTORequests_PrimitivesIgnored(t *testing.T) {
	ep := endpointInfo{bodyType: "String", returnType: "int"}
	ri := resourceInfo{imports: map[string]string{}}
	if reqs := collectDTORequests(ep, ri, "/abs", 0); len(reqs) != 0 {
		t.Fatalf("expected 0, got %d", len(reqs))
	}
}

func TestResolveSameFileClass(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "R.java")
	writeFile(t, dir, "R.java", `class R {} class UserDto { String name; }`)
	if got := resolveSameFileClass(p, "UserDto", dir); got != p {
		t.Fatalf("got %q", got)
	}
	if got := resolveSameFileClass(p, "Missing", dir); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveSamePackageClass(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto {}`)
	referrer := filepath.Join(dir, "R.java")
	if got := resolveSamePackageClass(referrer, "UserDto"); got != filepath.Join(dir, "UserDto.java") {
		t.Fatalf("got %q", got)
	}
	if got := resolveSamePackageClass(referrer, "Missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveImportPath(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/UserDto.java", `class UserDto {}`)
	got := resolveImportPath(dir, "com.example.UserDto")
	if got != filepath.Join(dir, "src/main/java/com/example/UserDto.java") {
		t.Fatalf("got %q", got)
	}
	if got := resolveImportPath(dir, "com.example.Missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveClassFieldsWithParams(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "UserDto.java")
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; private int age; }`)
	cache := map[string][]scanner.Field{}
	r, err := resolveClassFieldsWithParams(p, "UserDto", dir, cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.fields) != 2 {
		t.Fatalf("got %+v", r.fields)
	}
}

func TestResolveClassFields_Cached(t *testing.T) {
	cache := map[string][]scanner.Field{"X": {{Name: "a", Type: "string"}}}
	got, err := resolveClassFields("/ignored", "X", "/root", cache)
	if err != nil || len(got) != 1 {
		t.Fatalf("got %+v err %v", got, err)
	}
}

func TestResolveEnumFromAST(t *testing.T) {
	src := []byte(`enum Status { OPEN, CLOSED }`)
	root, _ := parseJava(src)
	cache := map[string][]scanner.Field{}
	fields := resolveEnumFromAST(root, src, "Status", cache)
	if len(fields) != 1 || len(fields[0].Enum) != 2 {
		t.Fatalf("got %+v", fields)
	}
}

func TestResolveEnumFromAST_NotFound(t *testing.T) {
	src := []byte(`enum Status { OPEN }`)
	root, _ := parseJava(src)
	if got := resolveEnumFromAST(root, src, "Other", map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestResolveClassFromASTWithParams(t *testing.T) {
	src := []byte(`class UserDto { private String name; }`)
	root, _ := parseJava(src)
	cache := map[string][]scanner.Field{}
	fields, _ := resolveClassFromASTWithParams(root, src, "UserDto", "/abs/UserDto.java", "/abs", cache)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}

func TestMergeParentFields(t *testing.T) {
	parent := []scanner.Field{{Name: "id"}, {Name: "name"}}
	own := []scanner.Field{{Name: "name"}, {Name: "email"}}
	got := mergeParentFields(parent, own)
	// id from parent + name,email from own (parent name skipped)
	if len(got) != 3 {
		t.Fatalf("got %d: %+v", len(got), got)
	}
	if got[0].Name != "id" {
		t.Fatalf("first should be id: %+v", got)
	}
}

func TestMergeParentFields_NoParent(t *testing.T) {
	own := []scanner.Field{{Name: "x"}}
	got := mergeParentFields(nil, own)
	if len(got) != 1 {
		t.Fatalf("got %+v", got)
	}
}

func TestResolveParentFields(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Base.java", `class Base { private Long id; }`)
	childSrc := `class Child extends Base { private String name; }`
	writeFile(t, dir, "Child.java", childSrc)
	root, _ := parseJava([]byte(childSrc))
	cls := findAllByType(root, "class_declaration")[0]
	fields := resolveParentFields(cls, []byte(childSrc), filepath.Join(dir, "Child.java"), dir, map[string]string{}, map[string][]scanner.Field{})
	if len(fields) != 1 || fields[0].Name != "id" {
		t.Fatalf("got %+v", fields)
	}
}

func TestResolveParentFields_NoSuper(t *testing.T) {
	src := []byte(`class C { String x; }`)
	root, _ := parseJava(src)
	cls := findAllByType(root, "class_declaration")[0]
	if got := resolveParentFields(cls, src, "/abs/C.java", "/abs", map[string]string{}, map[string][]scanner.Field{}); got != nil {
		t.Fatalf("expected nil, got %+v", got)
	}
}

func TestAssignDTOFields_Body(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{}}}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{isBody: true}, ep, fields)
	if len(ep.Request.Body.Fields) != 1 {
		t.Fatalf("got %+v", ep.Request.Body)
	}
}

func TestAssignDTOFields_Form(t *testing.T) {
	ep := &scanner.Endpoint{}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{isForm: true}, ep, fields)
	if ep.Request == nil || len(ep.Request.FormFields) != 1 {
		t.Fatalf("got %+v", ep.Request)
	}
}

func TestAssignDTOFields_Response(t *testing.T) {
	ep := &scanner.Endpoint{Responses: []scanner.Response{{Status: "200"}}}
	fields := []scanner.Field{{Name: "name", Type: "string"}}
	assignDTOFields(dtoRequest{}, ep, fields)
	if len(ep.Responses[0].Fields) != 1 {
		t.Fatalf("got %+v", ep.Responses)
	}
}

func TestResolveDTOTypeAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	dr := dtoRequest{
		typeName:    "UserDto",
		imports:     map[string]string{},
		referrer:    filepath.Join(dir, "Resource.java"),
		projectRoot: dir,
		epIdx:       0,
	}
	writeFile(t, dir, "Resource.java", `@Path("/x") class Resource {}`)
	cache := map[string][]scanner.Field{}
	fields := resolveDTOType(dr, dir, cache)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}

	// resolveAllDTOs path
	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, dir)
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("resolveAllDTOs did not assign: %+v", endpoints[0].Responses)
	}
}
