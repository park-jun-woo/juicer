//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestControllerVersion_Empty 테스트
package nestjs

import "testing"

func TestControllerVersion_Empty(t *testing.T) {
	src := []byte(`
import { Controller, Get } from '@nestjs/common';
@Controller('users')
export class UsersController {
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
	if ver != "" {
		t.Fatalf("expected empty, got %q", ver)
	}
}
