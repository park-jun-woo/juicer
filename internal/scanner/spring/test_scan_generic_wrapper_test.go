//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestScanGenericWrapper — 제네릭 래퍼 클래스 필드 추출 e2e 테스트
package spring

import (
	"testing"
)

func TestScanGenericWrapper(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/dto/PagedResponse.java", `
package com.example.dto;

import java.util.List;

public class PagedResponse<T> {
    private List<T> content;
    private int totalPages;
    private long totalElements;
}
`)
	writeFile(t, dir, "src/main/java/com/example/dto/AlbumResponse.java", `
package com.example.dto;

public class AlbumResponse {
    private Long id;
    private String title;
}
`)
	writeFile(t, dir, "src/main/java/com/example/controller/AlbumController.java", `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.dto.PagedResponse;
import com.example.dto.AlbumResponse;

@RestController
@RequestMapping("/api/albums")
public class AlbumController {

    @GetMapping
    public PagedResponse<AlbumResponse> list() {
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
	if len(ep.Responses) == 0 {
		t.Fatal("expected at least 1 response")
	}
	fields := ep.Responses[0].Fields
	if len(fields) == 0 {
		t.Fatal("expected response fields, got none")
	}
	found := map[string]string{}
	for _, f := range fields {
		found[f.Name] = f.Type
	}
	if found["totalPages"] != "integer" {
		t.Errorf("totalPages type = %q, want integer", found["totalPages"])
	}
	if found["totalElements"] != "integer" {
		t.Errorf("totalElements type = %q, want integer", found["totalElements"])
	}
	contentType := found["content"]
	if contentType != "[]AlbumResponse" {
		t.Errorf("content type = %q, want []AlbumResponse", contentType)
	}
}
