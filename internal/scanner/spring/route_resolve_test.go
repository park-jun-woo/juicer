//ff:func feature=scan type=test control=sequence topic=spring
//ff:what extractOneRoute / extractMethodEndpoints / extractMethodParams / buildControllerInfo / build*Endpoints / collectControllers / extractControllers / extractClassFields / extractEnumValues / extractResponseStatus / convertFieldTypes / fieldsToFormParams / ensureRequest / assignDTOFields / collectDTORequests / parseFile(All) / resolve* DTO 테스트
package spring

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func sFileInfo(t *testing.T, javaSrc string) *fileInfo {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{
		absPath:     "/abs/C.java",
		relPath:     "C.java",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		imports:     map[string]string{},
	}
}

const sampleController = `
@RestController
@RequestMapping("/users")
public class UserController {
    @GetMapping("/{id}")
    public UserDto get(@PathVariable("id") Long id) { return null; }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto create(@RequestBody CreateUserDto dto) { return null; }
}
`

func TestExtractOneRoute(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok || ep.method != "GET" || ep.path != "/{id}" || ep.handler != "get" {
		t.Fatalf("got %+v ok=%v", ep, ok)
	}
	if len(ep.params) != 1 {
		t.Fatalf("params: %+v", ep.params)
	}
}

func TestExtractOneRoute_NotRoute(t *testing.T) {
	fi := sFileInfo(t, `class C { public void helper() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	if _, ok := extractOneRoute(m, fi); ok {
		t.Fatal("expected false")
	}
}

func TestExtractMethodEndpoints(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}

func TestExtractMethodParams(t *testing.T) {
	fi := sFileInfo(t, `class C { void m(@PathVariable("id") Long id, @RequestParam("q") String q) {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractMethodParams(m, fi.src, ep, nil, "", "")
	if len(ep.params) != 1 || len(ep.query) != 1 {
		t.Fatalf("got %+v %+v", ep.params, ep.query)
	}
}

func TestBuildControllerInfo(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ci := buildControllerInfo(cls, fi)
	if ci.className != "UserController" || ci.prefix != "/users" {
		t.Fatalf("meta: %+v", ci)
	}
	if len(ci.endpoints) != 2 {
		t.Fatalf("endpoints: %+v", ci.endpoints)
	}
}

func TestExtractControllersAndCollect(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	got := collectControllers([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("collect: %d", len(got))
	}
}

func TestExtractControllers_NotController(t *testing.T) {
	fi := sFileInfo(t, `public class PlainService {}`)
	if c := extractControllers(fi); c != nil {
		t.Fatalf("expected nil, got %+v", c)
	}
}

func TestBuildAllEndpoints(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	controllers := extractControllers(fi)
	eps, _ := buildAllEndpoints(controllers, "/abs")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	found := false
	for _, e := range eps {
		if e.Path == "/users/{id}" {
			found = true
		}
	}
	if !found {
		t.Fatalf("missing /users/{id}: %+v", eps)
	}
}

func TestExtractClassFields(t *testing.T) {
	fi := sFileInfo(t, `class UserDto {
		@JsonProperty("user_name") private String name;
		private int age;
		private static String CONST;
	}`)
	cls := findAllByType(fi.root, "class_declaration")[0]
	fields := extractClassFields(cls, fi.src)
	if len(fields) != 2 {
		t.Fatalf("expected 2 (static skipped), got %d", len(fields))
	}
}

func TestExtractResponseStatus(t *testing.T) {
	fi := sFileInfo(t, `class C { @ResponseStatus(HttpStatus.CREATED) void m() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractResponseStatus(m, fi.src, ep)
	if ep.statusCode != "201" {
		t.Fatalf("got %q", ep.statusCode)
	}
}

func TestConvertFieldTypes(t *testing.T) {
	got := convertFieldTypes([]scanner.Field{{Name: "id", Type: "Long"}})
	if got[0].Type != "integer" {
		t.Fatalf("got %+v", got)
	}
}

func TestFieldsToFormParams(t *testing.T) {
	params := fieldsToFormParams([]scanner.Field{{Name: "a", Type: "string"}, {Name: "b", JSON: "bb", Type: "string"}})
	if len(params) != 2 || params[1].Name != "bb" {
		t.Fatalf("got %+v", params)
	}
}

func TestEnsureRequest(t *testing.T) {
	ep := &scanner.Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("expected request")
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

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ci := controllerInfo{imports: map[string]string{}}
	if reqs := collectDTORequests(ep, ci, "/abs", 0); len(reqs) != 2 {
		t.Fatalf("got %d", len(reqs))
	}
}

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserController.java", sampleController)
	fi, err := parseFile(dir, filepath.Join(dir, "UserController.java"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UserController.java" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UserController.java"),
		filepath.Join(dir, "missing.java"),
	})
	if len(files) != 1 {
		t.Fatalf("got %d", len(files))
	}
}

func TestResolveSameFileAndPackageClass(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "R.java")
	writeFile(t, dir, "R.java", `class R {} class UserDto { String name; }`)
	if resolveSameFileClass(p, "UserDto", dir) != p {
		t.Fatal("same file")
	}
	writeFile(t, dir, "Other.java", `class Other {}`)
	if resolveSamePackageClass(p, "Other") != filepath.Join(dir, "Other.java") {
		t.Fatal("same package")
	}
}

func TestResolveImportPath(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/UserDto.java", `class UserDto {}`)
	if resolveImportPath(dir, "com.example.UserDto") != filepath.Join(dir, "src/main/java/com/example/UserDto.java") {
		t.Fatal("import path")
	}
	if resolveImportPath(dir, "com.example.Missing") != "" {
		t.Fatal("missing import")
	}
}

func TestResolveClassFieldsWithParams(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "UserDto.java")
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	r, err := resolveClassFieldsWithParams(p, "UserDto", dir, map[string][]scanner.Field{})
	if err != nil || len(r.fields) != 1 {
		t.Fatalf("got %+v err %v", r, err)
	}
}

func TestResolveEnumFromAST(t *testing.T) {
	src := []byte(`enum Status { OPEN, CLOSED }`)
	root, _ := parseJava(src)
	fields := resolveEnumFromAST(root, src, "Status", map[string][]scanner.Field{})
	if len(fields) != 1 || len(fields[0].Enum) != 2 {
		t.Fatalf("got %+v", fields)
	}
}

func TestMergeParentFields(t *testing.T) {
	got := mergeParentFields([]scanner.Field{{Name: "id"}, {Name: "name"}}, []scanner.Field{{Name: "name"}})
	if len(got) != 2 || got[0].Name != "id" {
		t.Fatalf("got %+v", got)
	}
}

func TestResolveDTOTypeAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	writeFile(t, dir, "C.java", `@RestController class C {}`)
	dr := dtoRequest{
		typeName:    "UserDto",
		imports:     map[string]string{},
		referrer:    filepath.Join(dir, "C.java"),
		projectRoot: dir,
		epIdx:       0,
	}
	fields := resolveDTOType(dr, dir, map[string][]scanner.Field{})
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
	endpoints := []scanner.Endpoint{{Responses: []scanner.Response{{Status: "200"}}}}
	resolveAllDTOs([]dtoRequest{dr}, endpoints, dir)
	if len(endpoints[0].Responses[0].Fields) != 1 {
		t.Fatalf("resolveAllDTOs: %+v", endpoints[0].Responses)
	}
}
