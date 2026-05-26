//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_Basic 테스트
package nestjs

import "testing"

func TestExtractControllers_Basic(t *testing.T) {
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
	controllers := extractControllers(root, src, "test.ts", "test.ts")
	if len(controllers) != 1 {
		t.Fatalf("expected 1, got %d", len(controllers))
	}
}
