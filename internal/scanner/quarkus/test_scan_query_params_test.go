//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestScan_QueryParams -- @QueryParam + @DefaultValue 추출 테스트
package quarkus

import "testing"

func TestScan_QueryParams(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/ItemResource.java", `
package com.example.resource;

import javax.ws.rs.*;
import java.util.List;

@Path("/api/items")
public class ItemResource {

    @GET
    public List<String> list(
        @QueryParam("page") @DefaultValue("0") int page,
        @QueryParam("size") @DefaultValue("10") int size,
        @QueryParam("q") String q) {
        return null;
    }
}
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Method != "GET" {
		t.Errorf("method: want GET, got %s", ep.Method)
	}
	if ep.Path != "/api/items" {
		t.Errorf("path: want /api/items, got %s", ep.Path)
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.Query) != 3 {
		t.Fatalf("expected 3 query params, got %d", len(ep.Request.Query))
	}
	if ep.Request.Query[0].Name != "page" {
		t.Errorf("query[0] name: want page, got %s", ep.Request.Query[0].Name)
	}
	if ep.Request.Query[0].Default != "0" {
		t.Errorf("query[0] default: want 0, got %s", ep.Request.Query[0].Default)
	}
	if ep.Request.Query[1].Name != "size" {
		t.Errorf("query[1] name: want size, got %s", ep.Request.Query[1].Name)
	}
	if ep.Request.Query[1].Default != "10" {
		t.Errorf("query[1] default: want 10, got %s", ep.Request.Query[1].Default)
	}
	if ep.Request.Query[2].Name != "q" {
		t.Errorf("query[2] name: want q, got %s", ep.Request.Query[2].Name)
	}
}
