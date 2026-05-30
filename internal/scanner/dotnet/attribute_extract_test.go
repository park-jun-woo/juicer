//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what attribute/extract 함수 테스트
package dotnet

import (
	"path/filepath"
	"testing"
)

func TestAttributeFirstStringArg(t *testing.T) {
	root, src := parseCS(t, `class C { [Route("api/users")] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	if got := attributeFirstStringArg(attrs[0], src); got != "api/users" {
		t.Fatalf("got %q", got)
	}
}

func TestAttributeIntArgs(t *testing.T) {
	root, src := parseCS(t, `class C { [ProducesResponseType(201)] void m() {} }`)
	attrs := findAllByType(root, "attribute")
	if len(attrs) == 0 {
		t.Skip("no attribute")
	}
	ints := attributeIntArgs(attrs[0], src)
	if len(ints) != 1 || ints[0] != 201 {
		t.Fatalf("got %v", ints)
	}
}

func TestHasAttributeAndFind(t *testing.T) {
	root, src := parseCS(t, `[ApiController] class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !hasAttribute(cls, src, "ApiController") {
		t.Fatal("ApiController")
	}
	if hasAttribute(cls, src, "Missing") {
		t.Fatal("missing")
	}
	if findAttribute(cls, src, "ApiController") == nil {
		t.Fatal("find")
	}
}

func TestExtractHTTPMethodAndPath(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet("{id}")] public string Get(int id) { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	method, path, ok := extractHTTPMethodAndPath(m, src)
	if !ok || method != "GET" || path != "{id}" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}

func TestExtractHTTPMethodAndPath_None(t *testing.T) {
	root, src := parseCS(t, `class C { public string Helper() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if _, _, ok := extractHTTPMethodAndPath(m, src); ok {
		t.Fatal("expected false")
	}
}

func TestExtractClassRoute(t *testing.T) {
	root, src := parseCS(t, `[Route("api/[controller]")] class UsersController {}`)
	cls := findAllByType(root, "class_declaration")[0]
	got := extractClassRoute(cls, src, "Users")
	if got == "" {
		t.Fatal("expected non-empty route")
	}
}

func TestExtractParamNameAndType(t *testing.T) {
	root, src := parseCS(t, `class C { void m(int id) {} }`)
	params := findAllByType(root, "parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	if got := extractParamName(params[0], src); got != "id" {
		t.Fatalf("name: %q", got)
	}
	if got := extractParamType(params[0], src); got != "int" {
		t.Fatalf("type: %q", got)
	}
}

func TestExtractClassProps(t *testing.T) {
	root, src := parseCS(t, `class UserDto {
		public string Name { get; set; }
		public int Age { get; set; }
	}`)
	cls := findAllByType(root, "class_declaration")[0]
	props := extractClassProps(cls, src)
	if len(props) != 2 || props[0].Name != "Name" {
		t.Fatalf("got %+v", props)
	}
	if props[0].Type != "string" || props[1].Type != "integer" {
		t.Fatalf("types: %+v", props)
	}
}

func TestExtractOneProperty_Nullable(t *testing.T) {
	root, src := parseCS(t, `class C { public int? Age { get; set; } }`)
	props := findAllByType(root, "property_declaration")
	f := extractOneProperty(props[0], src)
	if f.Name != "Age" || !f.Nullable {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractUsings(t *testing.T) {
	root, src := parseCS(t, `using System;
using Microsoft.AspNetCore.Mvc;
namespace App {}`)
	usings := extractUsings(root, src)
	if len(usings) != 2 || usings[0] != "System" {
		t.Fatalf("got %v", usings)
	}
}

func TestIsApiController(t *testing.T) {
	root, src := parseCS(t, `[ApiController] class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if !isApiController(cls, src) {
		t.Fatal("ApiController attr")
	}
	root2, src2 := parseCS(t, `class D : ControllerBase {}`)
	cls2 := findAllByType(root2, "class_declaration")[0]
	if !isApiController(cls2, src2) {
		t.Fatal("ControllerBase base")
	}
	root3, src3 := parseCS(t, `class E {}`)
	cls3 := findAllByType(root3, "class_declaration")[0]
	if isApiController(cls3, src3) {
		t.Fatal("plain class")
	}
}

func TestFindCsFiles(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/UsersController.cs", "class C {}")
	writeFile(t, dir, "bin/Gen.cs", "class G {}")
	files, err := findCsFiles(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		rel, _ := filepath.Rel(dir, f)
		if rel != filepath.Join("Controllers", "UsersController.cs") {
			t.Errorf("unexpected: %s", rel)
		}
	}
}
