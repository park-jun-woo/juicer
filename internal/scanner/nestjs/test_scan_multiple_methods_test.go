//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what NestJS 다중 HTTP 메서드 스캔 테스트
package nestjs

import "testing"

func TestScan_MultipleHTTPMethods(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Get, Post, Put, Patch, Delete } from '@nestjs/common';

@Controller('items')
export class ItemsController {
  @Get()
  list() {}

  @Post()
  create() {}

  @Put(':id')
  update() {}

  @Patch(':id')
  partialUpdate() {}

  @Delete(':id')
  remove() {}
}
`
	writeFile(t, dir, "src/items.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 5 {
		t.Fatalf("expected 5 endpoints, got %d", len(result.Endpoints))
	}
	methods := map[string]bool{}
	for _, ep := range result.Endpoints {
		methods[ep.Method] = true
	}
	for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		if !methods[m] {
			t.Errorf("missing method: %s", m)
		}
	}
}
