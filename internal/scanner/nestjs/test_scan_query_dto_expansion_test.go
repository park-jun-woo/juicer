//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS @Query() DTO 필드 전개 스캔 테스트
package nestjs

import "testing"

func TestScan_QueryDTOExpansion(t *testing.T) {
	dir := t.TempDir()

	dto := `
export class PageOptionsDto {
  page: number;
  limit: number;
  q?: string;
}
`
	ctrl := `
import { Controller, Get, Query } from '@nestjs/common';
import { PageOptionsDto } from './dto/page-options.dto';

@Controller('users')
export class UsersController {
  @Get()
  findAll(@Query() dto: PageOptionsDto) {}
}
`
	writeFile(t, dir, "src/dto/page-options.dto.ts", dto)
	writeFile(t, dir, "src/users.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.Query) != 3 {
		t.Fatalf("expected 3 query params, got %d", len(ep.Request.Query))
	}
	q := ep.Request.Query
	if q[0].Name != "page" || q[0].Type != "number" {
		t.Errorf("param 0: want page/number, got %s/%s", q[0].Name, q[0].Type)
	}
	if q[1].Name != "limit" || q[1].Type != "number" {
		t.Errorf("param 1: want limit/number, got %s/%s", q[1].Name, q[1].Type)
	}
	if q[2].Name != "q" {
		t.Errorf("param 2: want q, got %s", q[2].Name)
	}
}
