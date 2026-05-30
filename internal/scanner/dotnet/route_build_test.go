//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what classify/build/extract route + controller 함수 테스트
package dotnet

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"

	sitter "github.com/smacker/go-tree-sitter"
)

func csFileInfo(t *testing.T, src string) *fileInfo {
	t.Helper()
	root, b := parseCS(t, src)
	return &fileInfo{
		absPath:     "/abs/C.cs",
		relPath:     "C.cs",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		usings:      extractUsings(root, b),
	}
}

const sampleCtrl = `
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase {
    [HttpGet("{id}")]
    public ActionResult<UserDto> Get(int id) { return Ok(); }

    [HttpPost]
    public ActionResult<UserDto> Create([FromBody] CreateUserDto dto) { return Ok(); }
}
`

func firstParamCS(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	root, b := parseCS(t, src)
	params := findAllByType(root, "parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	return params[0], b
}

func TestClassifyParam_FromBody(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromBody] UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if ep.bodyType != "UserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}

func TestClassifyParam_FromQuery(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromQuery] string q) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if len(ep.query) != 1 {
		t.Fatalf("got %+v", ep.query)
	}
}

func TestClassifyParam_FromRoute(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromRoute] int id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if len(ep.params) != 1 {
		t.Fatalf("got %+v", ep.params)
	}
}

func TestClassifyFormParam(t *testing.T) {
	ep := &endpointInfo{}
	classifyFormParam(ep, "IFormFile", "file")
	if len(ep.files) != 1 {
		t.Fatalf("file: %+v", ep.files)
	}
	ep2 := &endpointInfo{}
	classifyFormParam(ep2, "string", "name")
	if len(ep2.formFields) != 1 {
		t.Fatalf("form: %+v", ep2.formFields)
	}
}

func TestBuildEndpoint(t *testing.T) {
	ci := controllerInfo{prefix: "api/users", className: "UsersController", roles: []string{"Admin"}}
	ep := endpointInfo{method: "GET", path: "{id}", handler: "Get", params: []scanner.Param{{Name: "id", Type: "integer"}}, returnType: "UserDto"}
	got := buildEndpoint(ci, ep)
	if got.Method != "GET" || got.Path == "" {
		t.Fatalf("meta: %+v", got)
	}
	if got.Request == nil || len(got.Responses) != 1 {
		t.Fatalf("req/resp: %+v", got)
	}
}

func TestBuildRequest(t *testing.T) {
	if r := buildRequest(endpointInfo{}); r != nil {
		t.Fatalf("nil: %+v", r)
	}
	r := buildRequest(endpointInfo{bodyType: "UserDto"})
	if r == nil || r.Body == nil || r.Body.Method != "FromBody" {
		t.Fatalf("body: %+v", r)
	}
}

func TestBuildResponse(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("nil: %+v", r)
	}
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneRoute(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok || ep.method != "GET" || ep.handler != "Get" {
		t.Fatalf("got %+v ok=%v", ep, ok)
	}
}

func TestExtractMethodEndpoints(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}

func TestBuildControllerInfoAndExtract(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1, got %d", len(controllers))
	}
	if controllers[0].className != "UsersController" {
		t.Fatalf("name: %q", controllers[0].className)
	}
	if len(controllers[0].endpoints) != 2 {
		t.Fatalf("endpoints: %+v", controllers[0].endpoints)
	}
	got := collectControllers([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("collect: %d", len(got))
	}
}

func TestExtractAuthorizeRoles(t *testing.T) {
	root, src := parseCS(t, `class C { [Authorize(Roles = "Admin,User")] void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	roles := extractAuthorizeRoles(m, src)
	if len(roles) != 2 || roles[0] != "Admin" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractAuthorizeRoles_None(t *testing.T) {
	root, src := parseCS(t, `class C { void m() {} }`)
	m := findAllByType(root, "method_declaration")[0]
	if roles := extractAuthorizeRoles(m, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}

func TestApplyDataAnnotations(t *testing.T) {
	root, src := parseCS(t, `class C { [Required][StringLength(50)] public string Name { get; set; } }`)
	props := findAllByType(root, "property_declaration")
	f := &scanner.Field{}
	applyDataAnnotations(props[0], src, f)
	if f.Validate != "required" {
		t.Fatalf("validate: %q", f.Validate)
	}
	if f.MaxLength == nil || *f.MaxLength != 50 {
		t.Fatalf("maxlen: %v", f.MaxLength)
	}
}

func TestExpandRouteTokens(t *testing.T) {
	if got := expandRouteTokens("api/[controller]", "UsersController", ""); got != "api/users" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveControllerName(t *testing.T) {
	if resolveControllerName("UsersController") != "users" {
		t.Fatal("controller name")
	}
}

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UsersController.cs", sampleCtrl)
	fi, err := parseFile(dir, filepath.Join(dir, "UsersController.cs"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UsersController.cs" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UsersController.cs"),
		filepath.Join(dir, "missing.cs"),
	})
	if len(files) != 1 {
		t.Fatalf("got %d", len(files))
	}
}
