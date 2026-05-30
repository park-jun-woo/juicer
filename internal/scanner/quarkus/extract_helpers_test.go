//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what extractHTTPMethod / extractMethodPath / extractClassPath / extractParamName / extractParamType / extractDefaultValue / extractFieldName / extractFieldType / extractSuperclassName / extractTypeParams / firstStringArg / annotationIntValue / annotationElementValue / parseOneImport / extractImports / findJavaFiles 테스트
package quarkus

import (
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func parseQ(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func TestExtractHTTPMethod(t *testing.T) {
	root, src := parseQ(t, `class R { @GET public String list() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	method, ok := extractHTTPMethod(m, src)
	if !ok || method != "GET" {
		t.Fatalf("got %q %v", method, ok)
	}
}

func TestExtractHTTPMethod_None(t *testing.T) {
	root, src := parseQ(t, `class R { public String list() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if _, ok := extractHTTPMethod(m, src); ok {
		t.Fatal("expected false")
	}
}

func TestExtractMethodPath(t *testing.T) {
	root, src := parseQ(t, `class R { @GET @Path("/{id}") public String get() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if got := extractMethodPath(m, src); got != "/{id}" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractMethodPath_None(t *testing.T) {
	root, src := parseQ(t, `class R { @GET public String get() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if got := extractMethodPath(m, src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractClassPath(t *testing.T) {
	root, src := parseQ(t, `@Path("/users") class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPath(cls, src); got != "/users" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractParamNameAndType(t *testing.T) {
	root, src := parseQ(t, `class R { public void m(String name) {} }`)
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no params")
	}
	if got := extractParamName(params[0], src); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := extractParamType(params[0], src); got != "String" {
		t.Fatalf("type: %q", got)
	}
}

func TestExtractDefaultValue(t *testing.T) {
	root, src := parseQ(t, `class R { public void m(@DefaultValue("10") @QueryParam("limit") int limit) {} }`)
	params := findAllByType(root, "formal_parameter")
	if got := extractDefaultValue(params[0], src); got != "10" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFieldNameAndType(t *testing.T) {
	root, src := parseQ(t, `class D { private String name; }`)
	fields := findAllByType(root, "field_declaration")
	if got := extractFieldName(fields[0], src); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := extractFieldType(fields[0], src); got != "String" {
		t.Fatalf("type: %q", got)
	}
}

func TestExtractSuperclassName(t *testing.T) {
	root, src := parseQ(t, `class Child extends Base {}`)
	sc := findAllByType(root, "superclass")
	if len(sc) == 0 {
		t.Skip("no superclass")
	}
	if got := extractSuperclassName(sc[0], src); got != "Base" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractTypeParams(t *testing.T) {
	root, src := parseQ(t, `class R<T, U> {}`)
	cls := findAllByType(root, "class_declaration")[0]
	got := extractTypeParams(cls, src)
	if len(got) != 2 || got[0] != "T" {
		t.Fatalf("got %v", got)
	}
}

func TestExtractTypeParams_None(t *testing.T) {
	root, src := parseQ(t, `class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractTypeParams(cls, src); got != nil {
		t.Fatalf("got %v", got)
	}
}

func TestAnnotationIntValue(t *testing.T) {
	root, src := parseQ(t, `class C { @Size(min = 1, max = 10) String s; }`)
	ann := findAllByType(root, "annotation")[0]
	if v, ok := annotationIntValue(ann, src, "min"); !ok || v != 1 {
		t.Fatalf("min: %d %v", v, ok)
	}
	if _, ok := annotationIntValue(ann, src, "missing"); ok {
		t.Fatal("missing key")
	}
}

func TestAnnotationElementValue(t *testing.T) {
	root, src := parseQ(t, `class C { @Path(value = "/x") void m() {} }`)
	ann := findAllByType(root, "annotation")[0]
	if got := annotationElementValue(ann, src, "value"); got != "/x" {
		t.Fatalf("got %q", got)
	}
}

func TestParseOneImport(t *testing.T) {
	root, src := parseQ(t, `import com.example.UserDto;`)
	imps := findAllByType(root, "import_declaration")
	name, fqcn := parseOneImport(imps[0], src)
	if name != "UserDto" || fqcn != "com.example.UserDto" {
		t.Fatalf("got %q %q", name, fqcn)
	}
}

func TestParseOneImport_Static(t *testing.T) {
	root, src := parseQ(t, `import static com.example.Util.foo;`)
	imps := findAllByType(root, "import_declaration")
	name, _ := parseOneImport(imps[0], src)
	if name != "" {
		t.Fatalf("static import should yield empty, got %q", name)
	}
}

func TestParseOneImport_Wildcard(t *testing.T) {
	root, src := parseQ(t, `import com.example.*;`)
	imps := findAllByType(root, "import_declaration")
	name, _ := parseOneImport(imps[0], src)
	if name != "" {
		t.Fatalf("wildcard should yield empty, got %q", name)
	}
}

func TestExtractImports(t *testing.T) {
	root, src := parseQ(t, `import com.example.UserDto;
import java.util.List;
class R {}`)
	imports := extractImports(root, src)
	if imports["UserDto"] != "com.example.UserDto" {
		t.Fatalf("got %v", imports)
	}
}

func TestFirstStringArg(t *testing.T) {
	root, src := parseQ(t, `class C { @Path("/abc") void m() {} }`)
	ann := findAllByType(root, "annotation")
	if len(ann) == 0 {
		ann = findAllByType(root, "marker_annotation")
	}
	if got := firstStringArg(ann[0], src); got != "/abc" {
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
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("src", "Foo.java") {
			t.Errorf("unexpected: %s", rel)
		}
	}
}
