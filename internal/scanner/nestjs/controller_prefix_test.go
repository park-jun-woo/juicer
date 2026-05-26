//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestControllerPrefix_Found 테스트
package nestjs

import "testing"

func TestControllerPrefix_Found(t *testing.T) {
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
	prefix, ok := controllerPrefix(classes[0], src)
	if !ok {
		t.Fatal("expected controller prefix")
	}
	if prefix != "users" {
		t.Fatalf("expected users, got %q", prefix)
	}
}
