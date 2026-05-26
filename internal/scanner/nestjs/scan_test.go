//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 기본 컨트롤러 스캔 테스트
package nestjs

import "testing"

func TestScan_BasicController(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Get, Post, Param, Body, HttpCode } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';

@Controller('users')
export class UsersController {
  @Get()
  findAll() {}

  @Get(':id')
  findOne(@Param('id') id: string) {}

  @Post()
  @HttpCode(201)
  create(@Body() dto: CreateUserDto) {}
}
`
	writeFile(t, dir, "src/users.controller.ts", ctrl)

	dto := `
export class CreateUserDto {
  name: string;
  email: string;
  age?: number;
}
`
	writeFile(t, dir, "src/dto/create-user.dto.ts", dto)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" {
		t.Errorf("ep0 method: want GET, got %s", ep0.Method)
	}
	if ep0.Path != "/users" {
		t.Errorf("ep0 path: want /users, got %s", ep0.Path)
	}
	if ep0.Handler != "findAll" {
		t.Errorf("ep0 handler: want findAll, got %s", ep0.Handler)
	}

	ep1 := result.Endpoints[1]
	if ep1.Path != "/users/{id}" {
		t.Errorf("ep1 path: want /users/{id}, got %s", ep1.Path)
	}
	if ep1.Request == nil || len(ep1.Request.PathParams) != 1 {
		t.Fatalf("ep1 expected 1 path param")
	}
	if ep1.Request.PathParams[0].Name != "id" {
		t.Errorf("ep1 param name: want id, got %s", ep1.Request.PathParams[0].Name)
	}

	ep2 := result.Endpoints[2]
	if ep2.Method != "POST" {
		t.Errorf("ep2 method: want POST, got %s", ep2.Method)
	}
	if ep2.Request == nil || ep2.Request.Body == nil {
		t.Fatalf("ep2 expected body")
	}
	if ep2.Request.Body.TypeName != "CreateUserDto" {
		t.Errorf("ep2 body type: want CreateUserDto, got %s", ep2.Request.Body.TypeName)
	}
	if len(ep2.Request.Body.Fields) != 3 {
		t.Fatalf("ep2 body fields: want 3, got %d", len(ep2.Request.Body.Fields))
	}
	if ep2.Request.Body.Fields[0].Name != "name" {
		t.Errorf("field 0 name: want name, got %s", ep2.Request.Body.Fields[0].Name)
	}
	if len(ep2.Responses) == 0 || ep2.Responses[0].Status != "201" {
		t.Errorf("ep2 status: want 201")
	}
}
