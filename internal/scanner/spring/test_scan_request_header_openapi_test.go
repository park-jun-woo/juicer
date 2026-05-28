//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestScanRequestHeader_OpenAPI — @RequestHeader가 Request.Headers에 포함되는지 테스트
package spring

import (
	"testing"
)

func TestScanRequestHeader_OpenAPI(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/controller/TokenController.java", `
package com.example.controller;

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/tokens")
public class TokenController {

    @GetMapping("/validate")
    public void validate(@RequestHeader("Authorization") String auth) {}
}
`)
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	req := result.Endpoints[0].Request
	if req == nil {
		t.Fatal("expected request, got nil")
	}
	if len(req.Headers) != 1 {
		t.Fatalf("expected 1 header, got %d", len(req.Headers))
	}
	if req.Headers[0].Name != "Authorization" {
		t.Errorf("header name: want Authorization, got %s", req.Headers[0].Name)
	}
	if req.Headers[0].Type != "string" {
		t.Errorf("header type: want string, got %s", req.Headers[0].Type)
	}
}
