//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS 경로 변환 스캔 테스트
package nestjs

import "testing"

func TestScan_PathConversion(t *testing.T) {
	dir := t.TempDir()

	ctrl := `
import { Controller, Get } from '@nestjs/common';

@Controller('users')
export class UsersController {
  @Get(':userId/posts/:postId')
  getPost() {}
}
`
	writeFile(t, dir, "src/users.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	want := "/users/{userId}/posts/{postId}"
	if result.Endpoints[0].Path != want {
		t.Errorf("path: want %s, got %s", want, result.Endpoints[0].Path)
	}
}
