//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 쿼리 파라미터 스캔 테스트
package nestjs

import "testing"

func TestScan_QueryParams(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Get, Query } from '@nestjs/common';

@Controller('search')
export class SearchController {
  @Get()
  search(@Query('q') query: string, @Query('page') page: number) {}
}
`
	writeFile(t, dir, "src/search.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil || len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %v", ep.Request)
	}
	if ep.Request.Query[0].Name != "q" {
		t.Errorf("query 0 name: want q, got %s", ep.Request.Query[0].Name)
	}
	if ep.Request.Query[1].Name != "page" {
		t.Errorf("query 1 name: want page, got %s", ep.Request.Query[1].Name)
	}
}
