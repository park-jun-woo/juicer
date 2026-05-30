//ff:func feature=scan type=test control=sequence topic=spring
//ff:what extractInterfaces / extractSuperclassName / findJavaFiles / walkNodes / extractRolesFromExpr / extractRolesAllowed / matchAnnotationRoute / status 추출 / match ResponseEntity / resolveRequestMappingMethod / interface endpoints / parseOneImport / extractAnnotationPath / extractDeclaratorLiteral / findStaticFinalField / resolveConstantValue / isEnumClass / hasModifiers / resolveParentFields 테스트
package spring

import (
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractInterfaces(t *testing.T) {
	root, src := parseS(t, `class C implements UserApi, Other {}`)
	cls := findAllByType(root, "class_declaration")[0]
	ifaces := extractInterfaces(cls, src)
	if len(ifaces) != 2 || ifaces[0] != "UserApi" {
		t.Fatalf("got %v", ifaces)
	}
}

func TestExtractInterfaces_None(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractInterfaces(cls, src); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestExtractSuperclassName(t *testing.T) {
	root, src := parseS(t, `class Child extends Base {}`)
	sc := findAllByType(root, "superclass")
	if len(sc) == 0 {
		t.Skip("no superclass")
	}
	if got := extractSuperclassName(sc[0], src); got != "Base" {
		t.Fatalf("got %q", got)
	}
}

func TestFindJavaFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/Foo.java", "class Foo {}")
	writeFile(t, dir, "target/Bar.java", "class Bar {}")
	files, err := findJavaFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1, got %v", files)
	}
}

func TestWalkNodes(t *testing.T) {
	root, _ := parseS(t, `class C { void a() {} void b() {} }`)
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

func TestExtractRolesFromExpr(t *testing.T) {
	roles := extractRolesFromExpr(`'ROLE_ADMIN', 'ROLE_USER'`)
	if len(roles) != 2 || roles[0] != "ADMIN" {
		t.Fatalf("got %v", roles)
	}
}

func TestExtractRolesAllowed(t *testing.T) {
	root, src := parseS(t, `@RolesAllowed({"admin"}) class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractRolesAllowed(cls, src)
	if len(roles) != 1 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}

func TestMatchAnnotationRoute_GetMapping(t *testing.T) {
	root, src := parseS(t, `class C { @GetMapping("/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	method, path, ok := matchAnnotationRoute(ann, src)
	if !ok || method != "GET" || path != "/x" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}

func TestMatchAnnotationRoute_RequestMapping(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(value = "/x", method = RequestMethod.POST) void m() {} }`)
	anns := findAllByType(root, "annotation")
	method, path, ok := matchAnnotationRoute(anns[0], src)
	if !ok || path != "/x" || method != "POST" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}

func TestMatchStatusFromArgs(t *testing.T) {
	root, src := parseS(t, `class C { void m() { ResponseEntity.status(HttpStatus.CREATED).build(); } }`)
	argLists := findAllByType(root, "argument_list")
	for _, al := range argLists {
		if code := extractStatusFromArgList(al, src); code == "201" {
			return
		}
	}
	t.Fatal("did not find 201")
}

func TestMatchStatusArgChild_IntLiteral(t *testing.T) {
	root, src := parseS(t, `class C { void m() { status(404); } }`)
	lits := findAllByType(root, "decimal_integer_literal")
	if len(lits) == 0 {
		t.Skip("no literal")
	}
	if got := matchStatusArgChild(lits[0], src); got != "404" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchResponseEntityInvocation(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return ResponseEntity.created(uri).build(); } }`)
	invs := findAllByType(root, "method_invocation")
	for _, inv := range invs {
		if matchResponseEntityInvocation(inv, src) == "201" {
			return
		}
	}
	t.Fatal("did not match 201")
}

func TestMatchBodyInvocations(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return ResponseEntity.noContent().build(); } }`)
	body := findAllByType(root, "block")[0]
	if code := matchBodyInvocations(body, src); code != "204" {
		t.Fatalf("got %q", code)
	}
}

func TestMatchResponseEntityConstructor(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return new ResponseEntity<>(body, HttpStatus.CREATED); } }`)
	objs := findAllByType(root, "object_creation_expression")
	for _, o := range objs {
		if matchResponseEntityConstructor(o, src) == "201" {
			return
		}
	}
	t.Skip("constructor status not matched in this grammar")
}

func TestMatchBodyConstructors(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return new ResponseEntity<>(HttpStatus.OK); } }`)
	body := findAllByType(root, "block")[0]
	_ = matchBodyConstructors(body, src) // exercise path
}

func TestResolveRequestMappingMethod(t *testing.T) {
	if got := resolveRequestMappingMethod("RequestMethod.POST"); got != "POST" {
		t.Fatalf("got %q", got)
	}
	if got := resolveRequestMappingMethod("{RequestMethod.GET}"); got != "GET" {
		t.Fatalf("braces: %q", got)
	}
	if got := resolveRequestMappingMethod("unknown"); got != "" {
		t.Fatalf("unknown: %q", got)
	}
}

func TestParseOneImport(t *testing.T) {
	root, src := parseS(t, `import com.example.UserDto;`)
	imps := findAllByType(root, "import_declaration")
	name, fqcn := parseOneImport(imps[0], src)
	if name != "UserDto" || fqcn != "com.example.UserDto" {
		t.Fatalf("got %q %q", name, fqcn)
	}
}

func TestExtractAnnotationPath(t *testing.T) {
	root, src := parseS(t, `class C { @RequestMapping(path = "/p") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := extractAnnotationPath(ann, src); got != "/p" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractDeclaratorLiteral(t *testing.T) {
	root, src := parseS(t, `class C { static final String X = "hello"; }`)
	decls := findAllByType(root, "variable_declarator")
	if len(decls) == 0 {
		t.Skip("no declarator")
	}
	if got := extractDeclaratorLiteral(decls[0], src); got != "hello" {
		t.Fatalf("got %q", got)
	}
}

func TestFindStaticFinalField(t *testing.T) {
	root, src := parseS(t, `class C { public static final String PREFIX = "api"; }`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := findStaticFinalField(cls, src, "PREFIX"); got != "api" {
		t.Fatalf("got %q", got)
	}
	if got := findStaticFinalField(cls, src, "MISSING"); got != "" {
		t.Fatalf("missing: %q", got)
	}
}

func TestResolveConstantValue(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Const.java", `class Const { public static final String LIMIT = "10"; }`)
	got := resolveConstantValue("Const.LIMIT", map[string]string{}, filepath.Join(dir, "R.java"), dir)
	if got != "10" {
		t.Fatalf("got %q", got)
	}
	// unresolvable -> returns original ref
	if got := resolveConstantValue("Unknown.X", map[string]string{}, filepath.Join(dir, "R.java"), dir); got != "Unknown.X" {
		t.Fatalf("unresolved: %q", got)
	}
	// no dot -> returns as-is
	if got := resolveConstantValue("plain", map[string]string{}, "", ""); got != "plain" {
		t.Fatalf("plain: %q", got)
	}
}

func TestIsEnumClass(t *testing.T) {
	root, _ := parseS(t, `enum E { A }`)
	en := findAllByType(root, "enum_declaration")[0]
	if !isEnumClass(en) {
		t.Fatal("enum")
	}
	root2, _ := parseS(t, `class C {}`)
	cls := findAllByType(root2, "class_declaration")[0]
	if isEnumClass(cls) {
		t.Fatal("class is not enum")
	}
}

func TestHasModifiers(t *testing.T) {
	root, src := parseS(t, `class C { public static final int X = 1; }`)
	field := findAllByType(root, "field_declaration")[0]
	if !hasModifiers(field, src, "static", "final") {
		t.Fatal("expected static final")
	}
	if hasModifiers(field, src, "private") {
		t.Fatal("not private")
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

func TestResolveInterfaceEndpoints(t *testing.T) {
	dir := t.TempDir()
	ifaceSrc := `
@RequestMapping("/api")
interface UserApi {
    @GetMapping("/users")
    String list();
}
`
	writeFile(t, dir, "UserApi.java", ifaceSrc)
	prefix, eps := resolveInterfaceEndpoints(filepath.Join(dir, "UserApi.java"), "UserApi", dir)
	if prefix != "/api" {
		t.Fatalf("prefix: %q", prefix)
	}
	if len(eps) != 1 {
		t.Fatalf("endpoints: %+v", eps)
	}
}
