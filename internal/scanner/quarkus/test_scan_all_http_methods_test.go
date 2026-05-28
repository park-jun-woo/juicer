//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestScan_AllHTTPMethods -- GET/POST/PUT/DELETE/PATCH 전체 HTTP 메서드 테스트
package quarkus

import "testing"

func TestScan_AllHTTPMethods(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/CrudResource.java", `
package com.example.resource;

import javax.ws.rs.*;

@Path("/api/items")
public class CrudResource {

    @GET
    public String list() { return null; }

    @POST
    public String create() { return null; }

    @PUT
    @Path("/{id}")
    public String update(@PathParam("id") Long id) { return null; }

    @DELETE
    @Path("/{id}")
    public void delete(@PathParam("id") Long id) { }

    @PATCH
    @Path("/{id}")
    public String patch(@PathParam("id") Long id) { return null; }
}
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 5 {
		t.Fatalf("expected 5 endpoints, got %d", len(result.Endpoints))
	}

	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	for i, want := range methods {
		if result.Endpoints[i].Method != want {
			t.Errorf("ep[%d] method: want %s, got %s", i, want, result.Endpoints[i].Method)
		}
	}

	if result.Endpoints[0].Path != "/api/items" {
		t.Errorf("GET path: want /api/items, got %s", result.Endpoints[0].Path)
	}
	if result.Endpoints[2].Path != "/api/items/{id}" {
		t.Errorf("PUT path: want /api/items/{id}, got %s", result.Endpoints[2].Path)
	}
}
