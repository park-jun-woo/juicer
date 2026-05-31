//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractControllers_ClassAndMethodGuards 테스트
package nestjs

import "testing"

func TestExtractControllers_ClassAndMethodGuards(t *testing.T) {
	src := []byte(`
import { Controller, Get, UseGuards } from '@nestjs/common';
@Controller('app')
@UseGuards(JwtAuthGuard)
export class AppController {
  @UseGuards(RolesGuard)
  @Get('admin')
  admin() {}
}
`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	controllers := extractControllers(root, src, "test.ts", "test.ts", "/tmp")
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	ep := controllers[0].endpoints[0]
	if len(ep.middleware) != 2 {
		t.Fatalf("expected 2 middleware, got %d: %v", len(ep.middleware), ep.middleware)
	}
	if ep.middleware[0] != "JwtAuthGuard" {
		t.Fatalf("expected first=JwtAuthGuard, got %q", ep.middleware[0])
	}
	if ep.middleware[1] != "RolesGuard" {
		t.Fatalf("expected second=RolesGuard, got %q", ep.middleware[1])
	}
}
