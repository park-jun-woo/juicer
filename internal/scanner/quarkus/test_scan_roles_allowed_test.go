//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestScan_RolesAllowed -- @RolesAllowed 역할 추출 테스트
package quarkus

import "testing"

func TestScan_RolesAllowed(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "src/main/java/com/example/resource/AdminResource.java", `
package com.example.resource;

import javax.ws.rs.*;
import javax.annotation.security.RolesAllowed;

@Path("/api/admin")
@RolesAllowed("admin")
public class AdminResource {

    @GET
    public String dashboard() {
        return "ok";
    }

    @DELETE
    @Path("/{id}")
    @RolesAllowed("superadmin")
    public void deleteUser(@PathParam("id") Long id) {
    }
}
`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if len(ep0.Roles) != 1 || ep0.Roles[0] != "admin" {
		t.Errorf("ep0 roles: want [admin], got %v", ep0.Roles)
	}

	ep1 := result.Endpoints[1]
	if len(ep1.Roles) != 1 || ep1.Roles[0] != "superadmin" {
		t.Errorf("ep1 roles: want [superadmin], got %v", ep1.Roles)
	}
}
