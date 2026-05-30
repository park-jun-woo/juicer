//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what buildResourceInfo / buildResourceEndpoints / buildAllEndpoints / collectResources / extractResources / classifyBodyParam / extractMethodParams / matchResponseInvocation(s) / parseFile / parseAllFiles 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

const sampleResource = `
@Path("/users")
@RolesAllowed({"admin"})
public class UserResource {
    @GET
    @Path("/{id}")
    public UserDto get(@PathParam("id") Long id) { return null; }

    @POST
    public Response create(UserDto dto) { return Response.status(201).build(); }
}
`

func TestBuildResourceInfo(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ri := buildResourceInfo(cls, fi)
	if ri.className != "UserResource" || ri.prefix != "/users" {
		t.Fatalf("meta: %+v", ri)
	}
	if len(ri.roles) != 1 || len(ri.endpoints) != 2 {
		t.Fatalf("roles/endpoints: %+v", ri)
	}
}

func TestExtractResources(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	resources := extractResources(fi)
	if len(resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(resources))
	}
}

func TestExtractResources_NoPath(t *testing.T) {
	fi := qFileInfo(t, `public class PlainClass {}`)
	if r := extractResources(fi); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestCollectResources(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	got := collectResources([]*fileInfo{fi})
	if len(got) != 1 {
		t.Fatalf("got %d", len(got))
	}
}

func TestBuildResourceEndpointsAndAll(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	resources := extractResources(fi)
	eps, _ := buildAllEndpoints(resources, "/abs")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	// paths joined with class prefix
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

func TestClassifyBodyParam(t *testing.T) {
	ep := &endpointInfo{}
	classifyBodyParam("UserDto", "dto", ep)
	if ep.bodyType != "UserDto" || ep.bodyVarName != "dto" {
		t.Fatalf("got %+v", ep)
	}
	// primitive -> ignored
	ep2 := &endpointInfo{}
	classifyBodyParam("String", "s", ep2)
	if ep2.bodyType != "" {
		t.Fatalf("primitive should be ignored: %+v", ep2)
	}
	// second body ignored
	ep3 := &endpointInfo{bodyType: "First"}
	classifyBodyParam("Second", "x", ep3)
	if ep3.bodyType != "First" {
		t.Fatalf("second body should be ignored: %+v", ep3)
	}
}

func TestExtractMethodParams(t *testing.T) {
	fi := qFileInfo(t, `class R { void m(@PathParam("id") Long id, @QueryParam("q") String q) {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractMethodParams(m, fi.src, ep, nil, "", "")
	if len(ep.params) != 1 || len(ep.query) != 1 {
		t.Fatalf("got params=%+v query=%+v", ep.params, ep.query)
	}
}

func TestMatchResponseInvocation_Status(t *testing.T) {
	root, _ := parseJava([]byte(`class R { void m() { Response.status(404).build(); } }`))
	src := []byte(`class R { void m() { Response.status(404).build(); } }`)
	invs := findAllByType(root, "method_invocation")
	for _, inv := range invs {
		if code := matchResponseInvocation(inv, src); code == "404" {
			return
		}
	}
	t.Fatal("did not match 404")
}

func TestMatchResponseInvocations(t *testing.T) {
	root, _ := parseJava([]byte(`class R { void m() { return Response.status(204).build(); } }`))
	src := []byte(`class R { void m() { return Response.status(204).build(); } }`)
	body := findAllByType(root, "block")[0]
	if code := matchResponseInvocations(body, src); code != "204" {
		t.Fatalf("got %q", code)
	}
}

func TestParseFileAndAll(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserResource.java", sampleResource)
	fi, err := parseFile(dir, filepath.Join(dir, "UserResource.java"))
	if err != nil {
		t.Fatal(err)
	}
	if fi.relPath != "UserResource.java" {
		t.Fatalf("relPath: %q", fi.relPath)
	}
	files := parseAllFiles(dir, []string{
		filepath.Join(dir, "UserResource.java"),
		filepath.Join(dir, "missing.java"),
	})
	if len(files) != 1 {
		t.Fatalf("expected 1 parsed, got %d", len(files))
	}
}

func TestParseFile_Missing(t *testing.T) {
	if _, err := parseFile("/abs", "/no/such.java"); err == nil {
		t.Fatal("expected error")
	}
}
