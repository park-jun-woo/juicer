//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestControllerVersion_Found 테스트
package nestjs

import "testing"

func TestControllerVersion_Found(t *testing.T) {
	src := []byte(`
import { Controller, Get } from '@nestjs/common';
@Controller({ path: 'auth', version: '1' })
export class AuthController {
  @Get()
  findAll() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no classes found")
	}
	ver := controllerVersion(classes[0], src)
	if ver != "1" {
		t.Fatalf("expected 1, got %q", ver)
	}
}
