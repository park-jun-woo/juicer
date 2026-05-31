//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractControllers_ClassLevelGuards 테스트
package nestjs

import "testing"

func TestExtractControllers_ClassLevelGuards(t *testing.T) {
	src := []byte(`
import { Controller, Get, Post, UseGuards } from '@nestjs/common';
@Controller('tasks')
@UseGuards(JwtAuthGuard)
export class TasksController {
  @Get()
  findAll() {}

  @Post()
  create() {}
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
	ci := controllers[0]
	if len(ci.endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(ci.endpoints))
	}
	for _, ep := range ci.endpoints {
		if len(ep.middleware) != 1 || ep.middleware[0] != "JwtAuthGuard" {
			t.Fatalf("expected middleware=[JwtAuthGuard] for %s, got %v", ep.handler, ep.middleware)
		}
	}
}
