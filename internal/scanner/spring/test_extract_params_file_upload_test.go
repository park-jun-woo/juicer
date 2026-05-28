//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParams_FileUpload — @RequestPart 파일 업로드 추출 테스트
package spring

import "testing"

func TestExtractParams_FileUpload(t *testing.T) {
	src := []byte(`
package com.example;

import org.springframework.web.multipart.MultipartFile;

@RestController
public class TestController {

    @PostMapping("/upload")
    public void upload(@RequestPart("file") MultipartFile file) {}
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "TestController.java", absPath: "/test/TestController.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) != 1 || len(controllers[0].endpoints) != 1 {
		t.Fatalf("expected 1 controller with 1 endpoint")
	}
	ep := controllers[0].endpoints[0]
	if len(ep.files) != 1 {
		t.Fatalf("expected 1 file param, got %d", len(ep.files))
	}
	if ep.files[0].Name != "file" {
		t.Errorf("file param name: want file, got %s", ep.files[0].Name)
	}
}
