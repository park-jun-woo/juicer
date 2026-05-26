//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what NestJS URI versioning 스캔 통합 테스트
package nestjs

import "testing"

func TestScan_VersionedController(t *testing.T) {
	dir := t.TempDir()

	mainTS := `
import { NestFactory } from '@nestjs/core';
import { VersioningType } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.setGlobalPrefix('api');
  app.enableVersioning({ type: VersioningType.URI });
  await app.listen(3000);
}
bootstrap();
`
	writeFile(t, dir, "src/main.ts", mainTS)

	ctrl := `
import { Controller, Post } from '@nestjs/common';

@Controller({ path: 'auth', version: '1' })
export class AuthController {
  @Post('email/login')
  login() {}
}
`
	writeFile(t, dir, "src/auth.controller.ts", ctrl)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	if result.Endpoints[0].Path != "/api/v1/auth/email/login" {
		t.Errorf("path: want /api/v1/auth/email/login, got %s", result.Endpoints[0].Path)
	}
	if result.Endpoints[0].Method != "POST" {
		t.Errorf("method: want POST, got %s", result.Endpoints[0].Method)
	}
}
