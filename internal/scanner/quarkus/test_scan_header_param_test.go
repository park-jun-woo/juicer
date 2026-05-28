//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestScan_HeaderParam -- @HeaderParam 헤더 파라미터 추출 테스트
package quarkus

import "testing"

func TestScan_HeaderParam(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/ApiResource.java", `
package com.example.resource;

import javax.ws.rs.*;

@Path("/api/data")
public class ApiResource {

    @GET
    public String getData(@HeaderParam("X-Api-Key") String apiKey) {
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
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.Headers) != 1 {
		t.Fatalf("expected 1 header, got %d", len(ep.Request.Headers))
	}
	if ep.Request.Headers[0].Name != "X-Api-Key" {
		t.Errorf("header name: want X-Api-Key, got %s", ep.Request.Headers[0].Name)
	}
}
