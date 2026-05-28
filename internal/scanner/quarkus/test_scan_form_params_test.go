//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestScan_FormParams -- @FormParam 폼 파라미터 추출 테스트
package quarkus

import "testing"

func TestScan_FormParams(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/LoginResource.java", `
package com.example.resource;

import javax.ws.rs.*;

@Path("/auth")
public class LoginResource {

    @POST
    @Path("/login")
    public String login(
        @FormParam("username") String username,
        @FormParam("password") String password) {
        return "ok";
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
	if ep.Method != "POST" {
		t.Errorf("method: want POST, got %s", ep.Method)
	}
	if ep.Path != "/auth/login" {
		t.Errorf("path: want /auth/login, got %s", ep.Path)
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.FormFields) != 2 {
		t.Fatalf("expected 2 form fields, got %d", len(ep.Request.FormFields))
	}
	if ep.Request.FormFields[0].Name != "username" {
		t.Errorf("form[0] name: want username, got %s", ep.Request.FormFields[0].Name)
	}
	if ep.Request.FormFields[1].Name != "password" {
		t.Errorf("form[1] name: want password, got %s", ep.Request.FormFields[1].Name)
	}
}
