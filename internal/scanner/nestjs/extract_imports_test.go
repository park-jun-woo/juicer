//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractImports_Relative 테스트
package nestjs

import "testing"

func TestExtractImports_Relative(t *testing.T) {
	src := []byte(`
import { CreateUserDto } from './dto/create-user.dto';
import { Module } from '@nestjs/common';
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := extractImports(root, src)
	if result["CreateUserDto"] != "./dto/create-user.dto" {
		t.Fatalf("expected import path, got %v", result)
	}
	if _, ok := result["Module"]; ok {
		t.Fatal("should not include non-relative imports")
	}
}
